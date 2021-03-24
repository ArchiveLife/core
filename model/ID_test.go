package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateID(t *testing.T) {
	assert := assert.New(t)
	type args struct {
		entityType     string
		uriOrIDOrTitle interface{}
	}
	tests := []struct {
		name string
		args args
		want ID
	}{
		{
			"test create id with int",
			args{
				entityType:     "e1",
				uriOrIDOrTitle: 12,
			},
			"2d49eb6fad0f28d1f0d101dde91f97700630389787e1f65f3d1c68b342b3fac0",
		},
		{
			"test create id with string",
			args{
				entityType:     "e1",
				uriOrIDOrTitle: "12",
			},
			"2d49eb6fad0f28d1f0d101dde91f97700630389787e1f65f3d1c68b342b3fac0",
		},
		{
			"test create id with other type",
			args{
				entityType:     "e2",
				uriOrIDOrTitle: "12",
			},
			"a2690b1adbee34d474549807e93821bb4e31aea5e3ab7ba6c06fdaeeed280027",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CreateID(tt.args.entityType, tt.args.uriOrIDOrTitle)
			assert.Equal(tt.want, got, "CreateID() = %v, want %v", got, tt.want)
		})
	}
}
