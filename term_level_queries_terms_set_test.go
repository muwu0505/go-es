package goesdsl

import (
	"testing"
)

func TestQueryTermsSet_Valid(t *testing.T) {
	tests := []*TestBase{
		{
			name:  "valid query",
			query: NewTermsSetQuery("tags").SetTerms("search", "elasticsearch").SetMinimumShouldMatch(1),
			err:   false,
		},
		{
			name:  "missing field",
			query: NewTermsSetQuery("").SetTerms("search"),
			err:   true,
			want:  "field is required",
		},
		{
			name:  "missing value",
			query: NewTermsSetQuery("tags").SetMinimumShouldMatch(1),
			err:   true,
			want:  "terms is required",
		},
		{
			name:  "empty terms list",
			query: NewTermsSetQuery("tags").SetTerms().SetMinimumShouldMatch(1),
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
	}

	TestValid(t, tests)
}

func TestQueryTermsSet_Map(t *testing.T) {
	tests := []*TestBase{
		{
			name:  "basic query",
			query: NewTermsSetQuery("product_ids").SetTerms(1, 2, 3).SetMinimumShouldMatch(2),
			want: map[string]any{
				"terms_set": map[string]any{
					"product_ids": map[string]any{
						"terms":                []any{1, 2, 3},
						"minimum_should_match": 2,
					},
				},
			},
			err: false,
		},
		{
			name:  "with boost",
			query: NewTermsSetQuery("category").SetTerms("books", "ebooks").SetMinimumShouldMatchField("required_matches").SetBoost(1.5),
			want: map[string]any{
				"terms_set": map[string]any{
					"category": map[string]any{
						"terms":                      []any{"books", "ebooks"},
						"minimum_should_match_field": "required_matches",
						"boost":                      1.5,
					},
				},
			},
			err: false,
		},
		{
			name:  "with minimum should match script",
			query: NewTermsSetQuery("programming_languages").SetTerms("c++", "java", "php").SetMinimumShouldMatchScript(NewPainlessSourceScript("Math.min(params.num_terms, doc['required_matches'].value)", nil)),
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
			err: false,
		},
	}

	TestMap(t, tests)
}
