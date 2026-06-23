package goesdsl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestBase struct {
	name  string
	query IQuery
	err   bool
	want  any
}

func TestValid(t *testing.T, tests []*TestBase) {
	for _, query := range tests {
		t.Run(query.name, func(t *testing.T) {
			var q IQuery
			if query.query != nil {
				q = query.query
			}

			err := q.Valid()
			switch query.err {
			case true:
				assert.Error(t, err)
				assert.Equal(t, query.want, err.Error())
			case false:
				assert.NoError(t, err)
			}
		})
	}
}

func TestMap(t *testing.T, tests []*TestBase) {
	for _, query := range tests {
		t.Run(query.name, func(t *testing.T) {
			var q IQuery
			if query.query != nil {
				q = query.query
			}

			_map, err := q.ToMap()
			switch query.err {
			case true:
				assert.Error(t, err)
				assert.Nil(t, _map)
				assert.Equal(t, query.want.(string), err.Error())
			case false:
				assert.NoError(t, err)
				assert.Equal(t, query.want, _map)
			}
		})
	}
}
