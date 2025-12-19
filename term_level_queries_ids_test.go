package goesdsl

import (
	"testing"
)

func TestQueryIDs_Valid(t *testing.T) {
	tests := []*TestBase{
		{
			name:  "valid query",
			query: NewIDsQuery().SetValues("1", "2", "3"),
			err:   false,
		},
		{
			name:  "empty values",
			query: NewIDsQuery().SetValues(),
			err:   true,
			want:  "values is required",
		},
	}

	TestValid(t, tests)
}

func TestQueryIDs_Map(t *testing.T) {
	tests := []*TestBase{
		{
			name:  "basic query",
			query: NewIDsQuery().SetValues("100", "200"),
			want: map[string]any{
				"ids": map[string]any{
					"values": []string{"100", "200"},
				},
			},
			err: false,
		},
	}

	TestMap(t, tests)
}
