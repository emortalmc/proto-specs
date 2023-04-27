package kafkautils

import (
	"context"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

type ConsumerHandler interface {
	// Run starts the consumer handler
	// It will block until the context is cancelled
	Run(ctx context.Context)

	RegisterHandler(protoType proto.Message, consumer func(context.Context, *kafka.Message, proto.Message))
}

type consumerHandlerImpl struct {
	logger *zap.SugaredLogger
	reader *kafka.Reader

	handlers map[protoreflect.MessageType]func(context.Context, *kafka.Message, proto.Message)
}

func NewConsumerHandler(logger *zap.SugaredLogger,
	reader *kafka.Reader) ConsumerHandler {

	handler := &consumerHandlerImpl{
		logger: logger,
		reader: reader,

		handlers: make(map[protoreflect.MessageType]func(context.Context, *kafka.Message, proto.Message)),
	}

	return handler
}

func (c *consumerHandlerImpl) Run(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			c.logger.Infow("stopping consumer handler")
			return
		default:
			m, err := c.reader.ReadMessage(ctx)
			if err != nil {
				if err == context.Canceled {
					c.logger.Infow("stopping consumer handler")
					return
				}
				c.logger.Errorw("failed to read message", "error", err)
				continue
			}

			c.handleMessage(ctx, &m)
		}
	}
}

func (c *consumerHandlerImpl) handleMessage(ctx context.Context, m *kafka.Message) {
	protoTypeName, err := ProtoTypeFromHeaders(m.Headers)
	if err != nil {
		c.logger.Errorw("failed to parse proto type", "error", err, "offset", m.Offset, "headers", m.Headers)
		return
	}

	protoType, err := protoregistry.GlobalTypes.FindMessageByName(protoreflect.FullName(protoTypeName))
	if err != nil {
		c.logger.Errorw("failed to find proto type", "error", err, "offset", m.Offset, "headers", m.Headers)
		return
	}

	handler, ok := c.handlers[protoType]
	if !ok { // No handler registered for this proto type
		return
	}

	msg := protoType.New().Interface()
	if err := proto.Unmarshal(m.Value, msg); err != nil {
		c.logger.Errorw("failed to unmarshal message", "error", err, "offset", m.Offset, "headers", m.Headers)
		return
	}

	handler(ctx, m, msg)
}

func (c *consumerHandlerImpl) RegisterHandler(protoType proto.Message,
	consumer func(ctx context.Context, kafkaMsg *kafka.Message, msg proto.Message)) {

	c.handlers[protoType.ProtoReflect().Type()] = consumer
}
