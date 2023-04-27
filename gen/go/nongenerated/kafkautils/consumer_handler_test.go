package kafkautils

import (
	"context"
	"github.com/emortalmc/proto-specs/gen/go/message/common"
	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest"
	"google.golang.org/protobuf/proto"
	"testing"
)

func TestConsumerHandlerImpl_HandleMessage(t *testing.T) {
	testMessage := common.PlayerConnectMessage{
		PlayerId:       uuid.New().String(),
		PlayerUsername: "test",
		ServerId:       "test-server-wstg",
	}

	testMessageBytes, err := proto.Marshal(&testMessage)
	if err != nil {
		t.Fatalf("failed to marshal test message: %v", err)
	}

	tests := []struct {
		name string

		message       *kafka.Message
		handleMessage proto.Message

		shouldHandlerBeCalled bool
		wantErr               string
	}{
		{
			name: "success",
			message: &kafka.Message{
				Headers: []kafka.Header{{Key: "X-Proto-Type", Value: []byte(testMessage.ProtoReflect().Descriptor().FullName())}},
				Value:   testMessageBytes,
			},
			handleMessage: &testMessage,

			shouldHandlerBeCalled: true,
			wantErr:               "",
		},
		{
			name: "no proto type header",
			message: &kafka.Message{
				Value: testMessageBytes,
			},

			shouldHandlerBeCalled: false,
			wantErr:               "failed to parse proto type",
		},
		{
			name: "unknown proto type header",
			message: &kafka.Message{
				Headers: []kafka.Header{{Key: "X-Proto-Type", Value: []byte("invalid")}},
				Value:   testMessageBytes,
			},

			shouldHandlerBeCalled: false,
			wantErr:               "failed to find proto type",
		},
		{
			name: "no handler registered",
			message: &kafka.Message{
				Headers: []kafka.Header{{Key: "X-Proto-Type", Value: []byte(testMessage.ProtoReflect().Descriptor().FullName())}},
				Value:   testMessageBytes,
			},

			wantErr: "", // No error
		},
		{
			name: "invalid proto message",
			message: &kafka.Message{
				Headers: []kafka.Header{{Key: "X-Proto-Type", Value: []byte(testMessage.ProtoReflect().Descriptor().FullName())}},
				Value:   []byte("invalid"),
			},
			handleMessage: &testMessage,

			shouldHandlerBeCalled: false,
			wantErr:               "failed to unmarshal message",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger := zaptest.NewLogger(t, zaptest.WrapOptions(zap.Hooks(func(entry zapcore.Entry) error {
				if entry.Level == zapcore.ErrorLevel {
					if tt.wantErr == "" {
						t.Errorf("unexpected error: %v", entry.Message)
					} else if entry.Message != tt.wantErr {
						t.Errorf("expected error to be %v, got %v", tt.wantErr, entry.Message)
					}
				}
				return nil
			})))

			handler := NewConsumerHandler(logger.Sugar(), nil).(*consumerHandlerImpl)

			handlerCalled := false

			if tt.handleMessage != nil {
				handler.RegisterHandler(tt.handleMessage, func(ctx context.Context, kafkaMsg *kafka.Message, msg proto.Message) {
					handlerCalled = true

					if !proto.Equal(tt.handleMessage, msg) {
						t.Errorf("expected message to equal %v, got %v", tt.handleMessage, msg)
					}

					if !proto.Equal(tt.handleMessage, msg) {
						t.Errorf("expected message to equal %v, got %v", tt.handleMessage, msg)
					}
				})
			}

			handler.handleMessage(nil, tt.message)

			if tt.shouldHandlerBeCalled != handlerCalled {
				t.Errorf("expected handler to be called %v, got %v", tt.shouldHandlerBeCalled, handlerCalled)
			}

			if err := logger.Sync(); err != nil {
				t.Errorf("failed to sync logger: %v", err)
			}
		})
	}
}
