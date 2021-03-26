package message_type

// MessageType enum
type MessageType string

const (
	PlainText MessageType = "PlainText"
	RichText  MessageType = "RichText"
	Media     MessageType = "Media"
)
