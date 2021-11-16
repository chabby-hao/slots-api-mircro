FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux
#ENV GOPROXY https://goproxy.cn,direct

WORKDIR /build/zero

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
COPY ./etc /app/etc
RUN go build -ldflags="-s -w" -o /app/rmq ./rmq.go


FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
ENV TZ UTC

WORKDIR /app
COPY --from=builder /app/rmq /app/rmq
COPY --from=builder /app/etc /app/etc

CMD ["./rmq", "-f", "etc/rmq_prod.yaml"]
