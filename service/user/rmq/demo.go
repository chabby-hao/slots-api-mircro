package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sort"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/pkg/kmsg"
	"github.com/twmb/franz-go/pkg/sasl/aws"
)

var (
	seedBrokers = flag.String("brokers", "localhost:9092", "comma delimited list of seed brokers")
	topic       = flag.String("topic", "", "topic to consume from")
	style       = flag.String("commit-style", "autocommit", "commit style (which consume & commit is chosen); autocommit|records|uncommitted")
	group       = flag.String("group", "", "group to consume within")
	logger      = flag.Bool("logger", false, "if true, enable an info level logger")
)

func die(msg string, args ...interface{}) {
	fmt.Printf(msg+"\n", args...)
	os.Exit(1)
}

func main() {
	flag.Parse()

	styleNum := 0
	switch {
	case strings.HasPrefix("autocommit", *style):
	case strings.HasPrefix("records", *style):
		styleNum = 1
	case strings.HasPrefix("uncommitted", *style):
		styleNum = 2
	default:
		die("unrecognized style %s", *style)
	}

	opts := []kgo.Opt{
		kgo.SeedBrokers(strings.Split(*seedBrokers, ",")...),
		kgo.ConsumerGroup(*group),
		kgo.ConsumeTopics(*topic),

		kgo.SASL(aws.ManagedStreamingIAM(func(ctx context.Context) (aws.Auth, error) {

			sess, err := session.NewSession()
			if err != nil {
				die("unable to initialize aws session: %v", err)
			}

			val, err := sess.Config.Credentials.GetWithContext(ctx)
			if err != nil {
				return aws.Auth{}, err
			}
			return aws.Auth{
				AccessKey:    val.AccessKeyID,
				SecretKey:    val.SecretAccessKey,
				SessionToken: val.SessionToken,
				UserAgent:    "franz-go/creds_test/v1.0.0",
			}, nil
		})),

		kgo.Dialer((&tls.Dialer{NetDialer: &net.Dialer{Timeout: 10 * time.Second}}).DialContext),
	}
	if styleNum != 0 {
		opts = append(opts, kgo.DisableAutoCommit())
	}
	if *logger {
		opts = append(opts, kgo.WithLogger(kgo.BasicLogger(os.Stderr, kgo.LogLevelInfo, nil)))
	}
	cl, err := kgo.NewClient(opts...)
	if err != nil {
		die("unable to create client: %v", err)
	}
	defer cl.Close()

	resp, err := kmsg.NewPtrMetadataRequest().RequestWith(context.Background(), cl)
	if err != nil {
		die("unable to request metadata: %v", err)
	}

	if resp.ClusterID != nil {
		fmt.Printf("\nCLUSTER\n======\n%s\n", *resp.ClusterID)
	}

	fmt.Printf("\nBROKERS\n======\n")
	printBrokers(resp.ControllerID, resp.Brokers)

	fmt.Printf("\nTOPICS\n======\n")
	printTopics(resp.Topics)
	fmt.Println()

	go consume(cl, styleNum)

	sigs := make(chan os.Signal, 2)
	signal.Notify(sigs, os.Interrupt)

	<-sigs
	fmt.Println("received interrupt signal; closing client")
	done := make(chan struct{})
	go func() {
		defer close(done)
		cl.Close()
	}()

	select {
	case <-sigs:
		fmt.Println("received second interrupt signal; quitting without waiting for graceful close")
	case <-done:
	}
}

func consume(cl *kgo.Client, style int) {
	for {
		fetches := cl.PollFetches(context.Background())
		if fetches.IsClientClosed() {
			return
		}
		fetches.EachError(func(t string, p int32, err error) {
			die("fetch err topic %s partition %d: %v", t, p, err)
		})

		switch style {
		case 0:
			var seen int
			fetches.EachRecord(func(r *kgo.Record) {
				fmt.Println(r)
				seen++
			})
			fmt.Printf("processed %d records--autocommitting now allows the **prior** poll to be available for committing, nothing can be lost!\n", seen)

		case 1:
			var rs []*kgo.Record
			fetches.EachRecord(func(r *kgo.Record) {
				fmt.Println(r)
				rs = append(rs, r)
			})
			if err := cl.CommitRecords(context.Background(), rs...); err != nil {
				fmt.Printf("commit records failed: %v", err)
				continue
			}
			fmt.Printf("committed %d records individually--this demo does this in a naive way by just hanging on to all records, but you could just hang on to the max offset record per topic/partition!\n", len(rs))

		case 2:
			var seen int
			fetches.EachRecord(func(r *kgo.Record) {
				fmt.Println(r)
				seen++
			})
			if err := cl.CommitUncommittedOffsets(context.Background()); err != nil {
				fmt.Printf("commit records failed: %v", err)
				continue
			}
			fmt.Printf("committed %d records successfully--the recommended pattern, as followed in this demo, is to commit all uncommitted offsets after each poll!\n", seen)
		}
	}
}

func beginTabWrite() *tabwriter.Writer {
	return tabwriter.NewWriter(os.Stdout, 6, 4, 2, ' ', 0)
}

func printBrokers(controllerID int32, brokers []kmsg.MetadataResponseBroker) {
	sort.Slice(brokers, func(i, j int) bool {
		return brokers[i].NodeID < brokers[j].NodeID
	})

	tw := beginTabWrite()
	defer tw.Flush()

	fmt.Fprintf(tw, "ID\tHOST\tPORT\tRACK\n")
	for _, broker := range brokers {
		var controllerStar string
		if broker.NodeID == controllerID {
			controllerStar = "*"
		}

		var rack string
		if broker.Rack != nil {
			rack = *broker.Rack
		}

		fmt.Fprintf(tw, "%d%s\t%s\t%d\t%s\n",
			broker.NodeID, controllerStar, broker.Host, broker.Port, rack)
	}
}

func printTopics(topics []kmsg.MetadataResponseTopic) {
	// We request with no topic IDs, so we should not receive nil topics.
	sort.Slice(topics, func(i, j int) bool {
		return *topics[i].Topic < *topics[j].Topic
	})

	tw := beginTabWrite()
	defer tw.Flush()

	fmt.Fprintf(tw, "NAME\tPARTITIONS\tREPLICAS\n")
	for _, topic := range topics {
		parts := len(topic.Partitions)
		replicas := 0
		if parts > 0 {
			replicas = len(topic.Partitions[0].Replicas)
		}
		fmt.Fprintf(tw, "%s\t%d\t%d\n", *topic.Topic, parts, replicas)
	}
}
