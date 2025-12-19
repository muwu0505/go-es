package goesdsl

import (
	"testing"
)

func TestQueryExist_Valid(t *testing.T) {
	tests := []*TestBase{
		{
			name:  "valid query",
			query: NewExistQuery("user.id"),
			err:   false,
		},
		{
			name:  "missing field",
			query: NewExistQuery(""),
			err:   true,
			want:  "field is required",
		},
	}

	TestValid(t, tests)
}

func TestQueryExist_Map(t *testing.T) {
	tests := []*TestBase{
		{
			name:  "basic query",
			query: NewExistQuery("tags"),
			want: map[string]any{
				"exists": map[string]any{
					"field": "tags",
				},
			},
			err: false,
		},
	}

	TestMap(t, tests)
}
