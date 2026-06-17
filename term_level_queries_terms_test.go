package goesdsl

import (
	"testing"
)

func TestQueryTerms_Valid(t *testing.T) {
	tests := []*TestBase{
		{
			name:  "valid query",
			query: NewTermsQuery("tags").SetValue("search", "elasticsearch"),
			err:   false,
		},
		{
			name:  "missing field",
			query: NewTermsQuery("").SetValue("search"),
			err:   true,
			want:  "field is required",
		},
		{
			name:  "missing value",
			query: NewTermsQuery("tags"),
			err:   true,
			want:  "value is required",
		},
		{
			name:  "empty value list",
			query: NewTermsQuery("tags").SetValue(),
			err:   true,
			want:  "value is required",
		},
		{
			name: "missing lookup index",
			query: NewTermsQuery("tags").SetLookup(TermsLookup{
				ID:   "2",
				Path: "tags",
			}),
			err:  true,
			want: "lookup index is required",
		},
	}

	TestValid(t, tests)
}

func TestQueryTerms_Map(t *testing.T) {
	tests := []*TestBase{
		{
			name:  "basic query",
			query: NewTermsQuery("product_ids").SetValue(1, 2, 3),
			want: map[string]any{
				"terms": map[string]any{
					"product_ids": []any{1, 2, 3},
				},
			},
			err: false,
		},
		{
			name:  "with boost",
			query: NewTermsQuery("category").SetValue("books", "ebooks").SetBoost(1.5),
			want: map[string]any{
				"terms": map[string]any{
					"category": []any{"books", "ebooks"},
					"boost":    1.5,
				},
			},
			err: false,
		},
		{
			name: "with terms lookup",
			query: NewTermsQuery("color").SetLookup(TermsLookup{
				Index:   "my-index-000001",
				ID:      "2",
				Path:    "color",
				Routing: "user-1",
			}),
			want: map[string]any{
				"terms": map[string]any{
					"color": map[string]any{
						"index":   "my-index-000001",
						"id":      "2",
						"path":    "color",
						"routing": "user-1",
					},
				},
			},
			err: false,
		},
	}

	TestMap(t, tests)
}
