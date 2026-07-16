package goesdsl

import "testing"

// TestQueryTermsSet_Valid verifies terms set query validation.
func TestQueryTermsSet_Valid(t *testing.T) {
	tests := []*TestBase{
		{
			name:  "nil query",
			query: (*TermLevelQueriesTermsSet)(nil),
			err:   true,
			want:  "nil query",
		},
		{
			name:  "missing field",
			query: NewTermsSetQuery("").SetTerms("search"),
			err:   true,
			want:  "field is required",
		},
		{
			name:  "missing terms",
			query: NewTermsSetQuery("tags").SetMinimumShouldMatch(1),
			err:   true,
			want:  "terms is required",
		},
		{
			name:  "missing minimum should match",
			query: NewTermsSetQuery("tags").SetTerms("search"),
			err:   true,
			want:  "minimum_should_match is required",
		},
		{
			name: "multiple minimum should match options",
			query: &TermLevelQueriesTermsSet{
				field:                   "tags",
				terms:                   []any{"search"},
				minimumShouldMatch:      1,
				minimumShouldMatchField: "required_matches",
			},
			err:  true,
			want: "only one minimum_should_match option is allowed",
		},
		{
			name:  "full query",
			query: NewTermsSetQuery("category").SetTerms("books", "ebooks").SetMinimumShouldMatch(1).SetBoost(1.5),
			err:   false,
		},
		{
			name:  "minimum should match field",
			query: NewTermsSetQuery("category").SetTerms("books", "ebooks").SetMinimumShouldMatchField("required_matches"),
			err:   false,
		},
		{
			name:  "minimum should match script",
			query: NewTermsSetQuery("programming_languages").SetTerms("c++", "java", "php").SetMinimumShouldMatchScript(NewPainlessSourceScript("Math.min(params.num_terms, doc['required_matches'].value)", nil)),
			err:   false,
		},
	}

	TestValid(t, tests)
}

// TestQueryTermsSet_Json verifies terms set query JSON rendering.
func TestQueryTermsSet_Json(t *testing.T) {
	tests := []*TestBase{
		{
			name:  "full query",
			query: NewTermsSetQuery("category").SetTerms("books", "ebooks").SetMinimumShouldMatch(1).SetBoost(1.5),
			err:   false,
			want: map[string]any{
				"terms_set": map[string]any{
					"category": map[string]any{
						"terms":                []any{"books", "ebooks"},
						"minimum_should_match": 1,
						"boost":                1.5,
					},
				},
			},
		},
		{
			name:  "minimum should match field",
			query: NewTermsSetQuery("category").SetTerms("books", "ebooks").SetMinimumShouldMatchField("required_matches"),
			err:   false,
			want: map[string]any{
				"terms_set": map[string]any{
					"category": map[string]any{
						"terms":                      []any{"books", "ebooks"},
						"minimum_should_match_field": "required_matches",
					},
				},
			},
		},
		{
			name:  "minimum should match script",
			query: NewTermsSetQuery("programming_languages").SetTerms("c++", "java", "php").SetMinimumShouldMatchScript(NewPainlessSourceScript("Math.min(params.num_terms, doc['required_matches'].value)", nil)),
			err:   false,
			want: map[string]any{
				"terms_set": map[string]any{
					"programming_languages": map[string]any{
						"terms": []any{"c++", "java", "php"},
						"minimum_should_match_script": map[string]any{
							"script": map[string]any{
								"source": "Math.min(params.num_terms, doc['required_matches'].value)",
							},
						},
					},
				},
			},
		},
	}

	TestJson(t, tests)
}
