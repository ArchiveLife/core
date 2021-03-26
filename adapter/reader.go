package adapter

import (
	"github.com/ArchiveLife/core/model"
)

type ResourceReader interface {
	// call after options inject
	Init() error
	Next() (model.Resource, bool)
}
