package model

import "github.com/ArchiveLife/core/model/id"

type Author struct {
	ID            id.ID
	Age           *int
	FullName      string
	Address       *string
	Mobile        *string
	ExtAttributes map[string]interface{}
}
