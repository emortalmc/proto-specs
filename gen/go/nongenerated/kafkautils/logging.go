package kafkautils

import (
	"fmt"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
	"strings"
)

func CreateLogger(l *zap.SugaredLogger) kafka.Logger {
	return kafka.LoggerFunc(func(format string, args ...interface{}) {
		if strings.HasPrefix(format, "no messages received from kafka within the allocated time for partition") {
			return
		}

		l.Infow(fmt.Sprintf(format, args...))
	})
}

func CreateErrorLogger(l *zap.SugaredLogger) kafka.Logger {
	return kafka.LoggerFunc(func(format string, args ...interface{}) {
		l.Errorw(fmt.Sprintf(format, args...))
	})
}