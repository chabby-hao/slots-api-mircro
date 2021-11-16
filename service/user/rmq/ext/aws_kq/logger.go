package aws_kq

import (
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/twmb/franz-go/pkg/kgo"
)

type baseLogger struct {
	logger logx.Logger
}

func NewBaseLogger(logger logx.Logger) *baseLogger {
	return &baseLogger{
		logger: logger,
	}
}

func (b *baseLogger) Level() kgo.LogLevel {
	return kgo.LogLevelInfo
}
func (b *baseLogger) Log(level kgo.LogLevel, msg string, keyvals ...interface{}) {
	switch level {
	case kgo.LogLevelNone, kgo.LogLevelDebug, kgo.LogLevelInfo, kgo.LogLevelWarn:
		b.logger.Infof(msg, keyvals)
	case kgo.LogLevelError:
		b.logger.Errorf(msg, keyvals)
	}
}
