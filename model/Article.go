package model

import (
	"time"

	"github.com/ArchiveLife/core/model/id"
	"github.com/ArchiveLife/core/model/resource_type"
)

// Article
type Article struct {
	ID            id.ID
	Title         *string
	Author        *Author
	CoAuthors     []*Author
	DateTime      *time.Time
	Content       *string
	Tags          []string
	Type          string
	Medias        []*Media
	ExtAttributes map[string]interface{}
}

func (*Article) GetType() resource_type.ResourceType {
	return resource_type.Article
}
