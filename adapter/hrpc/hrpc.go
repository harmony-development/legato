package hrpc

import (
	goserver "github.com/harmony-development/hrpc/pkg/go-server"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// MarshalHRPC marshals a message to the hrpc protocol.
func MarshalHRPC(content goserver.VTProtoMessage, contentType string) ([]byte, error) {
	var response []byte

	var err error

	switch contentType {
	case "application/hrpc-json":
		response, err = protojson.Marshal(content)
	default:
		response, err = content.MarshalVT()
	}

	return response, err
}

func UnmarshalHRPC[T goserver.VTProtoMessage](content []byte, contentType string, messageType T) (T, error) {
	newMessage := proto.Clone(messageType).(T)

	var err error

	switch contentType {
	case "application/hrpc-json":
		if err := protojson.Unmarshal(content, newMessage); err != nil {
			return *new(T), err
		}
	default:
		if err := proto.Unmarshal(content, newMessage); err != nil {
			return *new(T), err
		}
	}
	return newMessage, err
}
