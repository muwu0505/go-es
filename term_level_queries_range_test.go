package goesdsl

import "testing"

// TestQueryRange_Valid verifies range query validation.
func TestQueryRange_Valid(t *testing.T) {
	tests := []*TestBase{
		{
			name:  "nil query",
			query: (*TermLevelQueriesRange)(nil),
			err:   true,
			want:  "nil query",
		},
		{
			name:  "missing field",
			query: NewRangeQuery("").SetGte(100).SetLte(200),
			err:   true,
			want:  "field is required",
		},
		{
			name:  "full query",
			query: NewRangeQuery("created_at").SetGt("2023-01-01").SetGte("2024-01-01").SetLt("2025-01-01").SetLte("2024-12-31").SetFormat("yyyy-MM-dd").SetTimeZone("+08:00").SetRelation(QueryRangeRelationWithin).SetBoost(2.0),
			err:   false,
		},
	}

	TestValid(t, tests)
}

// TestQueryRange_Json verifies range query JSON rendering.
func TestQueryRange_Json(t *testing.T) {
	tests := []*TestBase{
		{
			name:  "full query",
			query: NewRangeQuery("created_at").SetGt("2023-01-01").SetGte("2024-01-01").SetLt("2025-01-01").SetLte("2024-12-31").SetFormat("yyyy-MM-dd").SetTimeZone("+08:00").SetRelation(QueryRangeRelationWithin).SetBoost(2.0),
			err:   false,
			want: map[string]any{
				"range": map[string]any{
					"created_at": map[string]any{
						"gt":        "2023-01-01",
						"gte":       "2024-01-01",
						"lt":        "2025-01-01",
						"lte":       "2024-12-31",
						"format":    "yyyy-MM-dd",
						"time_zone": "+08:00",
						"relation":  QueryRangeRelationWithin,
						"boost":     2.0,
					},
				},
			},
		},
	}

	TestJson(t, tests)
}
