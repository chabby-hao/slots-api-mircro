package aws_kq

import (
	"context"
	"crypto/tls"
	"errors"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/tal-tech/go-zero/core/logx"
	_ "github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/proc"
	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/pkg/sasl/aws"
	"net"
	"time"
)

type (
	Consume = func(record *kgo.Record) error

	Consumer struct {
		client  *kgo.Client
		consume Consume
	}
)

func MustNewQueue(c AwsKqConf, consume Consume) *Consumer {
	q, err := NewQueue(c, consume)
	if err != nil {
		logx.Error(err)
		panic(err)
	}

	return q
}

func NewQueue(c AwsKqConf, consume Consume) (*Consumer, error) {
	if err := c.SetUp(); err != nil {
		return nil, err
	}

	opts := []kgo.Opt{
		kgo.SeedBrokers(c.Brokers...),
		kgo.ConsumeTopics(c.Topic),
		kgo.ConsumerGroup(c.Group),
		kgo.SASL(aws.ManagedStreamingIAM(func(ctx context.Context) (aws.Auth, error) {
			sess, err := session.NewSession()
			if err != nil {
				return aws.Auth{}, err
			}

			val, err := sess.Config.Credentials.GetWithContext(ctx)
			if err != nil {
				return aws.Auth{}, err
			}

			return aws.Auth{
				AccessKey:    val.AccessKeyID,
				SecretKey:    val.SecretAccessKey,
				SessionToken: val.SessionToken,
			}, nil
		})),
		kgo.Dialer((&tls.Dialer{
			NetDialer: &net.Dialer{
				Timeout: 10 * time.Second,
			},
		}).DialContext),
		//kgo.WithLogger(kgo.BasicLogger(os.Stderr, kgo.LogLevelInfo, nil)),
		kgo.WithLogger(NewBaseLogger(logx.WithContext(context.Background()))),
		kgo.DisableAutoCommit(),
	}

	cl, err := kgo.NewClient(opts...)
	if err != nil {
		return nil, err
	}

	q := &Consumer{
		client:  cl,
		consume: consume,
	}

	proc.AddShutdownListener(func() {
		cl.Close()
	})

	return q, nil
}

func (kf *Consumer) Start() {
	for {
		fetches := kf.client.PollRecords(context.Background(), 100)
		if fetches.IsClientClosed() {
			err := errors.New("client is closed")
			logx.Error(err)
			panic(err)
		}
		fetches.EachError(func(t string, p int32, err error) {
			logx.Error(err)
			panic(err)
		})
		recordsCount := len(fetches.Records())
		logx.Infof("got records: %d", recordsCount)

		var rs []*kgo.Record
		fetches.EachRecord(func(record *kgo.Record) {
			err := kf.consume(record)
			if err == nil {
				rs = append(rs, record)
			} else {
				logx.Error(err)
			}
		})

		if len(rs) > 0 {
			err := kf.client.CommitRecords(context.Background(), rs...)
			if err != nil {
				logx.Error(err)
			} else {
				ciCount := len(rs)
				logx.Infof("committed records: %d, failed: %d", ciCount, recordsCount-ciCount)
			}
		}
	}
}

func (kf *Consumer) Stop() {
	kf.client.Close()
}
