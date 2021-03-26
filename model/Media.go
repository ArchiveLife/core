package model

import (
	"net/url"

	"github.com/ArchiveLife/core/model/id"
)

type Media struct {
	ID       id.ID
	MimeType *string
	// the local or remote url
	URL *url.URL
}
