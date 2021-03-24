package adapter

import "github.com/ArchiveLife/model/model"

// ArticleConsumer function, make client could consume large stream data
type ArticleConsumer func(*model.Article)

// ArchiveService
type ArchiveService interface {
	// get the name of service
	GetName() string
	// get the description of service
	GetDescription() string
	// get the options of service
	GetOptions() []*Option
	// run service with input options
	Run(ArticleConsumer, ...OptionValue)
}
