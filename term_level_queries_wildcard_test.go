package goesdsl

import "testing"

// TestQueryWildcard_Valid verifies wildcard query validation.
func TestQueryWildcard_Valid(t *testing.T) {
	tests := []*TestBase{
		{
			name:  "nil query",
			query: (*TermLevelQueriesWildcard)(nil),
			err:   true,
			want:  "nil query",
		},
		{
			name:  "missing field",
			query: NewWildcardQuery("").SetValue("search"),
			err:   true,
			want:  "field is required",
		},
		{
			name:  "missing value",
			query: NewWildcardQuery("tags"),
			err:   true,
			want:  "value is required",
		},
		{
			name:  "full query",
			query: NewWildcardQuery("tags").SetWildcard("search*").SetBoost(1.5).SetCaseInsensitive(true).SetRewrite(QueriesRewriteParameterConstantScoreBlended),
			err:   false,
		},
	}

	TestValid(t, tests)
}

// TestQueryWildcard_Json verifies wildcard query JSON rendering.
func TestQueryWildcard_Json(t *testing.T) {
	tests := []*TestBase{
		{
			name:  "full query",
			query: NewWildcardQuery("tags").SetWildcard("search*").SetBoost(1.5).SetCaseInsensitive(true).SetRewrite(QueriesRewriteParameterConstantScoreBlended),
			err:   false,
			want: map[string]any{
				"wildcard": map[string]any{
					"tags": map[string]any{
						"wildcard":         "search*",
						"boost":            1.5,
						"case_insensitive": true,
						"rewrite":          QueriesRewriteParameterConstantScoreBlended,
					},
				},
			},
		},
	}

	TestJson(t, tests)
}
