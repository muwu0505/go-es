package goesdsl

import (
	"testing"
)

func TestQueryTermsSet_Valid(t *testing.T) {
	tests := []*TestBase{
		{
			name:  "valid query",
			query: NewTermsSetQuery("tags").SetTerms("search", "elasticsearch"),
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
			query: NewTermsSetQuery("tags"),
			err:   true,
			want:  "terms is required",
		},
	}

	TestValid(t, tests)
}

func TestQueryTermsSet_Map(t *testing.T) {
	tests := []*TestBase{
		{
			name:  "basic query",
			query: NewTermsSetQuery("product_ids").SetTerms(1, 2, 3),
			want: map[string]any{
				"terms_set": map[string]any{
					"product_ids": map[string]any{
						"terms": []any{1, 2, 3},
					},
				},
			},
			err: false,
		},
		{
			name:  "with boost",
			query: NewTermsSetQuery("category").SetTerms("books", "ebooks").SetMinimumShouldMatch(1.5),
			want: map[string]any{
				"terms_set": map[string]any{
					"category": map[string]any{
						"terms":                []any{"books", "ebooks"},
						"minimum_should_match": 1.5,
					},
				},
			},
			err: false,
		},
	}

	TestMap(t, tests)
}
