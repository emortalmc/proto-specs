package kafkautils

import (
	"fmt"
	"github.com/segmentio/kafka-go"
)

var ErrNoProtoTypeHeader = fmt.Errorf("no proto type found in message headers")

func ProtoTypeFromHeaders(headers []kafka.Header) (string, error) {
	for _, header := range headers {
		if header.Key == "X-Proto-Type" {
			return string(header.Value), nil
		}
	}
	return "", ErrNoProtoTypeHeader
}
