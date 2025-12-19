package goesdsl

import (
	"testing"
)

func TestQueryPrefix_Valid(t *testing.T) {
	tests := []*TestBase{
		{
			name:  "valid query",
			query: NewPrefixQuery("name").SetValue("john"),
			err:   false,
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

func TestQueryPrefix_Map(t *testing.T) {
	tests := []*TestBase{
		{
			name:  "basic query",
			query: NewPrefixQuery("product.code").SetValue("XJ"),
			want: map[string]any{
				"prefix": map[string]any{
					"product.code": map[string]any{
						"value": "XJ",
					},
				},
			},
			err: false,
		},
		{
			name:  "full options",
			query: NewPrefixQuery("title").SetValue("elastic").SetRewrite(QueriesRewriteParameterScoringBoolean).SetCaseInsensitive(true),
			want: map[string]any{
				"prefix": map[string]any{
					"title": map[string]any{
						"value":            "elastic",
						"rewrite":          "scoring_boolean",
						"case_insensitive": true,
					},
				},
			},
			err: false,
		},
	}

	TestMap(t, tests)
}
