package serial

import "github.com/golang/protobuf/proto"

// ToTypedMessage converts a proto Message into TypedMessage.
func ToTypedMessage(message proto.Message) *TypedMessage {
	if message == nil {
		return nil
	}
	settings, _ := proto.Marshal(message)
	return &TypedMessage{
		Type:  GetMessageType(message),
		Value: settings,
	}
}

// GetMessageType returns the name of this proto Message.
func GetMessageType(message proto.Message) string {
	return proto.MessageName(message)
}