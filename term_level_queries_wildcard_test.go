package goesdsl

import (
	"testing"
)

func TestQueryWildcard_Valid(t *testing.T) {
	tests := []*TestBase{
		{
			name:  "valid query",
			query: NewWildcardQuery("tags").SetValue("search"),
			err:   false,
		},
		{
			name:  "missing field",
			query: NewWildcardQuery("").SetValue("search"),
			err:   true,
			want:  "field is required",
		},
		{
			name:  "nil value",
			query: NewWildcardQuery("tags"),
			err:   true,
			want:  "value is required",
		},
	}

	TestValid(t, tests)
}

func TestQueryWildcard_Map(t *testing.T) {
	tests := []*TestBase{
		{
			name:  "basic query",
			query: NewWildcardQuery("tags").SetValue("search"),
			want: map[string]any{
				"wildcard": map[string]any{
					"tags": map[string]any{
						"value": "search",
					},
				},
			},
			err: false,
		},
		{
			name:  "with boost",
			query: NewWildcardQuery("tags").SetWildcard("search*").SetBoost(1.5).SetCaseInsensitive(true).SetRewrite(QueriesRewriteParameterConstantScoreBlended),
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
			err: false,
		},
	}

	TestMap(t, tests)
}
