package goesdsl

import "testing"

// TestQueryIDs_Valid verifies IDs query validation.
func TestQueryIDs_Valid(t *testing.T) {
	tests := []*TestBase{
		{
			name:  "nil query",
			query: (*TermLevelQueriesIDs)(nil),
			err:   true,
			want:  "nil query",
		},
		{
			name:  "missing values",
			query: NewIDsQuery().SetValues(),
			err:   true,
			want:  "values is required",
		},
		{
			name:  "full query",
			query: NewIDsQuery().SetValues("100", "200"),
			err:   false,
		},
	}

	TestValid(t, tests)
}

// TestQueryIDs_Json verifies IDs query JSON rendering.
func TestQueryIDs_Json(t *testing.T) {
	tests := []*TestBase{
		{
			name:  "full query",
			query: NewIDsQuery().SetValues("100", "200"),
			err:   false,
			want: map[string]any{
				"ids": map[string]any{
					"values": []string{"100", "200"},
				},
			},
		},
	}

	TestJson(t, tests)
}
