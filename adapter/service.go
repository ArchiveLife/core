package adapter

import (
	"github.com/ArchiveLife/core/model"
)

// ResourceConsumer function, make client could consume large stream data
type ResourceConsumer func(model.Resource)

// ArchiveService
type ArchiveService interface {
	// get the name of service
	GetName() string
	// get the description of service
	GetDescription() string
	// get the options of service
	GetOptions() []*Option
	// run service with input options
	Run(ResourceConsumer, ...*OptionValue) error
}

// NewServiceWrapper instance
func NewServiceWrapper(name, description string, reader ResourceReader, options ...*Option) *GenericServiceWrapper {
	return &GenericServiceWrapper{
		name,
		description,
		options,
		reader,
	}
}

type GenericServiceWrapper struct {
	name        string
	description string
	options     []*Option
	reader      ResourceReader
}

func (s *GenericServiceWrapper) GetName() string {
	return s.name
}

func (s *GenericServiceWrapper) GetDescription() string {
	return s.description
}

func (s *GenericServiceWrapper) GetOptions() []*Option {
	return s.options
}

// Run with dynamic options (blocking)
func (s *GenericServiceWrapper) Run(consumer ResourceConsumer, argOptValues ...*OptionValue) error {
	FillInstanceByValues(s.reader, argOptValues...)

	if err := s.reader.Init(); err != nil {
		return err
	}

	for {
		article, next := s.reader.Next()
		if article != nil {
			consumer(article)
		}
		if !next {
			break
		}
	}

	return nil

}
