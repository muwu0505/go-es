package goesdsl

import (
	"testing"
)

func TestQueryTerm_Valid(t *testing.T) {
	tests := []*TestBase{
		{
			name:  "valid query",
			query: NewTermQuery("status").SetValue("active"),
			err:   false,
			want:  nil,
		},
		{
			name:  "missing field",
			query: NewTermQuery("").SetValue("active"),
			err:   true,
			want:  "field is required",
		},
		{
			name:  "missing value",
			query: NewTermQuery("status"),
			err:   true,
			want:  "value is required",
		},
	}

	TestValid(t, tests)
}

func TestQueryTerm_Map(t *testing.T) {
	tests := []*TestBase{
		{
			name:  "basic query",
			query: NewTermQuery("status.keyword").SetValue("published"),
			want: map[string]any{
				"term": map[string]any{
					"status.keyword": map[string]any{
						"value": "published",
					},
				},
			},
			err: false,
		},
		{
			name:  "with boost and case_insensitive",
			query: NewTermQuery("role").SetValue("Admin").SetBoost(2.0).SetCaseInsensitive(true),
			want: map[string]any{
				"term": map[string]any{
					"role": map[string]any{
						"value":            "Admin",
						"boost":            2.0,
						"case_insensitive": true,
					},
				},
			},
			err: false,
		},
	}

	TestMap(t, tests)
}
