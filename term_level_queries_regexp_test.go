package goesdsl

import (
	"testing"
)

func TestQueryRegexp_Valid(t *testing.T) {
	tests := []*TestBase{
		{
			name:  "valid query",
			query: NewRegexpQuery("name").SetValue("k.*y"),
			err:   false,
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
	}

	TestValid(t, tests)
}

func TestQueryRegexp_Map(t *testing.T) {
	tests := []*TestBase{
		{
			name:  "basic query",
			query: NewRegexpQuery("product.code").SetValue("X[0-9]+"),
			want: map[string]any{
				"regexp": map[string]any{
					"product.code": map[string]any{
						"value": "X[0-9]+",
					},
				},
			},
			err: false,
		},
		{
			name:  "full options",
			query: NewRegexpQuery("title").SetValue("elasti.*").SetFlags(QueriesFlagsAll).SetCaseInsensitive(true).SetMaxDeterminizedStates(20000).SetRewrite(QueriesRewriteParameterConstantScoreBlended),
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
			err: false,
		},
	}

	TestMap(t, tests)
}
