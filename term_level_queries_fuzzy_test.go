package goesdsl

import "testing"

// TestQueryFuzzy_Valid verifies fuzzy query validation.
func TestQueryFuzzy_Valid(t *testing.T) {
	tests := []*TestBase{
		{
			name:  "nil query",
			query: (*TermLevelQueriesFuzzy)(nil),
			err:   true,
			want:  "nil query",
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
		{
			name: "full query",
			query: NewFuzzyQuery("name").SetValue("john").SetFuzziness(QueriesFuzzinessAuto).
				SetMaxExpansions(100).SetPrefixLength(2).SetTranspositions(true).SetRewrite(QueriesRewriteParameterConstantScore),
			err: false,
		},
	}

	TestValid(t, tests)
}

// TestQueryFuzzy_Json verifies fuzzy query JSON rendering.
func TestQueryFuzzy_Json(t *testing.T) {
	tests := []*TestBase{
		{
			name: "full query",
			query: NewFuzzyQuery("name").SetValue("john").SetFuzziness(QueriesFuzzinessAuto).
				SetMaxExpansions(100).SetPrefixLength(2).SetTranspositions(true).SetRewrite(QueriesRewriteParameterConstantScore),
			err: false,
			want: map[string]any{
				"fuzzy": map[string]any{
					"name": map[string]any{
						"value":          "john",
						"fuzziness":      QueriesFuzzinessAuto,
						"max_expansions": int32(100),
						"prefix_length":  int32(2),
						"transpositions": true,
						"rewrite":        QueriesRewriteParameterConstantScore,
					},
				},
			},
		},
	}

	TestJson(t, tests)
}
