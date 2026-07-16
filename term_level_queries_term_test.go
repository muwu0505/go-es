package goesdsl

import "testing"

// TestQueryTerm_Valid verifies term query validation.
func TestQueryTerm_Valid(t *testing.T) {
	tests := []*TestBase{
		{
			name:  "nil query",
			query: (*TermLevelQueriesTerm)(nil),
			err:   true,
			want:  "nil query",
		},
		{
			name:  "missing field",
			query: NewTermQuery("").SetValue("active"),
			err:   true,
			want:  "field is required",
		},
		{
			name:  "missing value",
			query: NewTermQuery("status"),
			err:   true,
			want:  "value is required",
		},
		{
			name:  "full query",
			query: NewTermQuery("role").SetValue("Admin").SetBoost(2.0).SetCaseInsensitive(true),
			err:   false,
		},
	}

	TestValid(t, tests)
}

// TestQueryTerm_Json verifies term query JSON rendering.
func TestQueryTerm_Json(t *testing.T) {
	tests := []*TestBase{
		{
			name:  "full query",
			query: NewTermQuery("role").SetValue("Admin").SetBoost(2.0).SetCaseInsensitive(true),
			err:   false,
			want: map[string]any{
				"term": map[string]any{
					"role": map[string]any{
						"value":            "Admin",
						"boost":            2.0,
						"case_insensitive": true,
					},
				},
			},
		},
	}

	TestJson(t, tests)
}
