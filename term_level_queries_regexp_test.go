package goesdsl

import "testing"

// TestQueryRegexp_Valid verifies regexp query validation.
func TestQueryRegexp_Valid(t *testing.T) {
	tests := []*TestBase{
		{
			name:  "nil query",
			query: (*TermLevelQueriesRegexp)(nil),
			err:   true,
			want:  "nil query",
		},
		{
			name:  "missing field",
			query: NewRegexpQuery("").SetValue("abc.*"),
			err:   true,
			want:  "field is required",
		},
		{
			name:  "missing value",
			query: NewRegexpQuery("name"),
			err:   true,
			want:  "value is required",
		},
		{
			name:  "full query",
			query: NewRegexpQuery("title").SetValue("elasti.*").SetFlags(QueriesFlagsAll).SetCaseInsensitive(true).SetMaxDeterminizedStates(20000).SetRewrite(QueriesRewriteParameterConstantScoreBlended),
			err:   false,
		},
	}

	TestValid(t, tests)
}

// TestQueryRegexp_Json verifies regexp query JSON rendering.
func TestQueryRegexp_Json(t *testing.T) {
	tests := []*TestBase{
		{
			name:  "full query",
			query: NewRegexpQuery("title").SetValue("elasti.*").SetFlags(QueriesFlagsAll).SetCaseInsensitive(true).SetMaxDeterminizedStates(20000).SetRewrite(QueriesRewriteParameterConstantScoreBlended),
			err:   false,
			want: map[string]any{
				"regexp": map[string]any{
					"title": map[string]any{
						"value":                   "elasti.*",
						"flags":                   QueriesFlagsAll,
						"case_insensitive":        true,
						"max_determinized_states": int64(20000),
						"rewrite":                 QueriesRewriteParameterConstantScoreBlended,
					},
				},
			},
		},
	}

	TestJson(t, tests)
}
