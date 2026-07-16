package goesdsl

import "testing"

// TestQueryPrefix_Valid verifies prefix query validation.
func TestQueryPrefix_Valid(t *testing.T) {
	tests := []*TestBase{
		{
			name:  "nil query",
			query: (*TermLevelQueriesPrefix)(nil),
			err:   true,
			want:  "nil query",
		},
		{
			name:  "missing field",
			query: NewPrefixQuery("").SetValue("test"),
			err:   true,
			want:  "field is required",
		},
		{
			name:  "missing value",
			query: NewPrefixQuery("name"),
			err:   true,
			want:  "value is required",
		},
	}

	TestValid(t, tests)
}

// TestQueryPrefix_Json verifies prefix query JSON rendering.
func TestQueryPrefix_Json(t *testing.T) {
	tests := []*TestBase{
		{
			name:  "full query",
			query: NewPrefixQuery("title").SetValue("elastic").SetRewrite(QueriesRewriteParameterScoringBoolean).SetCaseInsensitive(true),
			err:   false,
			want: map[string]any{
				"prefix": map[string]any{
					"title": map[string]any{
						"value":            "elastic",
						"rewrite":          QueriesRewriteParameterScoringBoolean,
						"case_insensitive": true,
					},
				},
			},
		},
	}

	TestJson(t, tests)
}
