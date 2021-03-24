package adapter

import "github.com/ArchiveLife/model/model"

type ArchiveService interface {
	GetOptions() []*Option
	Run(...OptionValue) []*model.Article
}
