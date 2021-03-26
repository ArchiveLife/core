package model

import (
	"time"

	"github.com/ArchiveLife/core/model/id"
	"github.com/ArchiveLife/core/model/message_type"
	"github.com/ArchiveLife/core/model/resource_type"
)

type Message struct {
	ID       id.ID
	DateTime *time.Time
	From     *Author
	To       *Author
	Content  *string
	Media    []*Media
	Type     message_type.MessageType
}

func (*Message) GetType() resource_type.ResourceType {
	return resource_type.Message
}
