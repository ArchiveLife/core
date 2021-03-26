package model

import "github.com/ArchiveLife/core/model/resource_type"

// Resource abstract interface of ArchiveLife
type Resource interface {
	GetType() resource_type.ResourceType
}
