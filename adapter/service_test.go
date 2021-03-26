package adapter

import (
	"testing"

	"github.com/ArchiveLife/core/model"
	"github.com/ArchiveLife/core/model/resource_type"
	"github.com/stretchr/testify/assert"
)

type service01Reader struct {
	cur int
	Max int
}

func (s *service01Reader) Init() error {
	s.cur = 0
	return nil
}

func (s *service01Reader) Next() (model.Resource, bool) {
	s.cur++
	if s.cur > s.Max {
		return nil, false
	}
	return &model.Article{}, true
}

func TestNewServiceWrapper(t *testing.T) {

	assert := assert.New(t)
	type args struct {
		name        string
		description string
		reader      ResourceReader
		options     []*Option
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"baisc test with wrapper",
			args{
				"service01",
				"test service",
				&service01Reader{},
				[]*Option{
					{Order: 0, Name: "Max"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewServiceWrapper(tt.args.name, tt.args.description, tt.args.reader, tt.args.options...)
			assert.NotNil(got)
			count := 0
			err := got.Run(func(a model.Resource) {
				count++
				assert.Equal(resource_type.Article, a.GetType())
				arti, ok := a.(*model.Article)
				assert.True(ok)
				assert.NotNil(arti)
			}, &OptionValue{Option: Option{Name: "Max"}, Value: 10})
			assert.Nil(err)
			assert.Equal(10, count)
		})
	}
}
