package goesdsl

import (
	"testing"
)

func TestQueryRange_Valid(t *testing.T) {
	tests := []*TestBase{
		{
			name:  "valid query",
			query: NewRangeQuery("age").SetGt(18).SetLte(65),
			err:   false,
		},
		{
			name:  "missing field",
			query: NewRangeQuery("").SetGte(100).SetLte(200),
			err:   true,
			want:  "field is required",
		},
	}

	TestValid(t, tests)
}

func TestQueryRange_Map(t *testing.T) {
	tests := []*TestBase{
		{
			name:  "basic range",
			query: NewRangeQuery("price").SetGt(10.0).SetLte(100.0),
			want: map[string]any{
				"range": map[string]any{
					"price": map[string]any{
						"gt":  10.0,
						"lte": 100.0,
					},
				},
			},
			err: false,
		},
		{
			name:  "date range with options",
			query: NewRangeQuery("created_at").SetGte("2024-01-01").SetLt("2024-12-31").SetFormat("yyyy-MM-dd").SetTimeZone("+08:00").SetRelation(QueryRangeRelationWithin).SetBoost(2.0),
			want: map[string]any{
				"range": map[string]any{
					"created_at": map[string]any{
						"gte":       "2024-01-01",
						"lt":        "2024-12-31",
						"format":    "yyyy-MM-dd",
						"time_zone": "+08:00",
						"relation":  QueryRangeRelationWithin,
						"boost":     2.0,
					},
				},
			},
			err: false,
		},
		{
			name:  "only gte",
			query: NewRangeQuery("score").SetGte(90),
			want: map[string]any{
				"range": map[string]any{
					"score": map[string]any{
						"gte": 90,
					},
				},
			},
			err: false,
		},
	}

	TestMap(t, tests)
}
