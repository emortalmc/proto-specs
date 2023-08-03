package gameutils

import (
	"errors"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

var ErrElementNotFound = errors.New("element not found")

func UnmarshalElement(elements []*anypb.Any, v proto.Message) (int, error) {
	msgType := string(v.ProtoReflect().Descriptor().FullName())

	for i, element := range elements {
		if element.TypeUrl == msgType {
			return i, element.UnmarshalTo(v)
		}
	}

	return -1, ErrElementNotFound
}

func ElementsContains(elements []*anypb.Any, v proto.Message) bool {
	msgType := string(v.ProtoReflect().Descriptor().FullName())

	for _, element := range elements {
		if element.TypeUrl == msgType {
			return true
		}
	}

	return false
}
