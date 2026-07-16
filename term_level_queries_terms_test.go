package goesdsl

import "testing"

// TestQueryTerms_Valid verifies terms query validation.
func TestQueryTerms_Valid(t *testing.T) {
	tests := []*TestBase{
		{
			name:  "nil query",
			query: (*TermLevelQueriesTerms)(nil),
			err:   true,
			want:  "nil query",
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
			name: "missing lookup index",
			query: NewTermsQuery("tags").SetLookup(TermsLookup{
				ID:   "2",
				Path: "tags",
			}),
			err:  true,
			want: "lookup index is required",
		},
		{
			name: "missing lookup id",
			query: NewTermsQuery("tags").SetLookup(TermsLookup{
				Index: "my-index-000001",
				Path:  "tags",
			}),
			err:  true,
			want: "lookup id is required",
		},
		{
			name: "missing lookup path",
			query: NewTermsQuery("tags").SetLookup(TermsLookup{
				Index: "my-index-000001",
				ID:    "2",
			}),
			err:  true,
			want: "lookup path is required",
		},
		{
			name:  "full query",
			query: NewTermsQuery("category").SetValue("books", "ebooks").SetBoost(1.5),
			err:   false,
		},
		{
			name: "lookup query",
			query: NewTermsQuery("color").SetLookup(TermsLookup{
				Index:   "my-index-000001",
				ID:      "2",
				Path:    "color",
				Routing: "user-1",
			}),
			err: false,
		},
	}

	TestValid(t, tests)
}

// TestQueryTerms_Json verifies terms query JSON rendering.
func TestQueryTerms_Json(t *testing.T) {
	tests := []*TestBase{
		{
			name:  "full query",
			query: NewTermsQuery("category").SetValue("books", "ebooks").SetBoost(1.5),
			err:   false,
			want: map[string]any{
				"terms": map[string]any{
					"category": []any{"books", "ebooks"},
					"boost":    1.5,
				},
			},
		},
		{
			name: "lookup query",
			query: NewTermsQuery("color").SetLookup(TermsLookup{
				Index:   "my-index-000001",
				ID:      "2",
				Path:    "color",
				Routing: "user-1",
			}),
			err: false,
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
		},
	}

	TestJson(t, tests)
}
