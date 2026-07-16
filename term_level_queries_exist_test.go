package goesdsl

import "testing"

// TestQueryExist_Valid verifies exists query validation.
func TestQueryExist_Valid(t *testing.T) {
	tests := []*TestBase{
		{
			name:  "nil query",
			query: (*TermLevelQueriesExist)(nil),
			err:   true,
			want:  "nil query",
		},
		{
			name:  "missing field",
			query: NewExistQuery(""),
			err:   true,
			want:  "field is required",
		},
		{
			name:  "full query",
			query: NewExistQuery("tags"),
			err:   false,
		},
	}

	TestValid(t, tests)
}

// TestQueryExist_Json verifies exists query JSON rendering.
func TestQueryExist_Json(t *testing.T) {
	tests := []*TestBase{
		{
			name:  "full query",
			query: NewExistQuery("tags"),
			err:   false,
			want: map[string]any{
				"exists": map[string]any{
					"field": "tags",
				},
			},
		},
	}

	TestJson(t, tests)
}
