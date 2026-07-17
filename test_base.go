package goesdsl

import (
	"encoding/json"
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
				if !assert.Error(t, err) {
					return
				}
				assert.Equal(t, query.want, err.Error())
			case false:
				assert.NoError(t, err)
			}
		})
	}
}

func TestJson(t *testing.T, tests []*TestBase) {
	for _, query := range tests {
		t.Run(query.name, func(t *testing.T) {
			if query.err {
				t.Fatalf("TestJson only accepts success cases; move %q to TestValid", query.name)
			}

			var q IQuery
			if query.query != nil {
				q = query.query
			}

			got, err := q.MarshalJson()
			assert.NoError(t, err)

			want, err := json.Marshal(query.want)
			assert.NoError(t, err)
			assert.JSONEq(t, string(want), got)
		})
	}
}
