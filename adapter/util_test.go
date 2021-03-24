package adapter

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateInstanceByValues(t *testing.T) {

	type TestSt01 struct {
		Name string
		Age  int
	}
	r := func(options TestSt01) {

	}
	assert := assert.New(t)
	type args struct {
		t      reflect.Type
		values []*OptionValue
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"test fill option with values",
			args{
				reflect.TypeOf(r).In(0),
				[]*OptionValue{
					{Option{Name: "Name"}, "Theo"},
					{Option{Name: "Age"}, 16},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CreateInstanceByValues(tt.args.t, tt.args.values...)
			assert.NotNil(got)
			value := got.Interface().(TestSt01)
			assert.Equal("Theo", value.Name)
			assert.Equal(16, value.Age)
		})
	}
}
