package kafkautils

import (
	"github.com/segmentio/kafka-go"
	"testing"
)

func TestProtoTypeFromHeaders(t *testing.T) {
	tests := []struct {
		name    string
		headers []kafka.Header
		want    string
		wantErr error
	}{
		{
			name: "proto type found",
			headers: []kafka.Header{
				{
					Key:   "X-Proto-Type",
					Value: []byte("test"),
				},
			},
			want:    "test",
		},
		{
			name: "proto type not found",
			headers: []kafka.Header{
				{
					Key: "Random",
					Value: []byte("test"),
				},
			},
			wantErr: ErrNoProtoTypeHeader,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ProtoTypeFromHeaders(tt.headers)
			if err != tt.wantErr {
				t.Errorf("ProtoTypeFromHeaders() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ProtoTypeFromHeaders() got = %v, want %v", got, tt.want)
			}
		})
	}
}
