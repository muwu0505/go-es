package goesdsl

import (
	"testing"
)

func TestQueryFuzzy_Valid(t *testing.T) {
	tests := []*TestBase{
		{
			name:  "valid query",
			query: NewFuzzyQuery("title").SetValue("elastic"),
			err:   false,
			want:  nil,
		},
		{
			name:  "missing field",
			query: NewFuzzyQuery("").SetValue("elastic"),
			err:   true,
			want:  "field is required",
		},
		{
			name:  "missing value",
			query: NewFuzzyQuery("title"),
			err:   true,
			want:  "value is required",
		},
	}

	TestValid(t, tests)
}

func TestQueryFuzzy_Map(t *testing.T) {
	tests := []*TestBase{
		{
			name:  "basic query",
			query: NewFuzzyQuery("content").SetValue("golang"),
			err:   false,
			want: map[string]any{
				"fuzzy": map[string]any{
					"content": map[string]any{
						"value": "golang",
					},
				},
			},
		},
		{
			name:  "full options",
			query: NewFuzzyQuery("name").SetValue("john").SetFuzziness("AUTO").SetMaxExpansions(100).SetPrefixLength(2).SetTranspositions(true).SetRewrite(QueriesRewriteParameterConstantScore),
			err:   false,
			want: map[string]any{
				"fuzzy": map[string]any{
					"name": map[string]any{
						"value":          "john",
						"fuzziness":      "AUTO",
						"max_expansions": int32(100),
						"prefix_length":  int32(2),
						"transpositions": true,
						"rewrite":        "constant_score",
					},
				},
			},
		},
	}

	TestMap(t, tests)
}
