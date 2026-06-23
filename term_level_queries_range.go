package goesdsl

import (
	"errors"
)

/*
	doc: https://www.elastic.co/docs/reference/query-languages/query-dsl/query-dsl-range-query
*/

type TermLevelQueriesRange struct {
	field    string
	gt       any
	gte      any
	lt       any
	lte      any
	format   string
	relation string
	timeZone string
	boost    *float64
}

const (
	QueryRangeRelationIntersects = "INTERSECTS" // default
	QueryRangeRelationContains   = "CONTAINS"
	QueryRangeRelationWithin     = "WITHIN"
)

func NewRangeQuery(field string) *TermLevelQueriesRange {
	return NewTermLevelQueriesRange(field)
}

func (q *TermLevelQueriesRange) Valid() error {
	if q == nil {
		return errors.New("nil query")
	}

	if q.field == "" {
		return errors.New("field is required")
	}

	return nil
}

func (q *TermLevelQueriesRange) ToMap() (map[string]any, error) {
	err := q.Valid()
	if err != nil {
		return nil, err
	}

	rangeQuery := make(map[string]any)

	if q.gt != nil {
		rangeQuery["gt"] = q.gt
	}
	if q.gte != nil {
		rangeQuery["gte"] = q.gte
	}
	if q.lt != nil {
		rangeQuery["lt"] = q.lt
	}
	if q.lte != nil {
		rangeQuery["lte"] = q.lte
	}
	if q.format != "" {
		rangeQuery["format"] = q.format
	}
	if q.relation != "" {
		rangeQuery["relation"] = q.relation
	}
	if q.timeZone != "" {
		rangeQuery["time_zone"] = q.timeZone
	}
	if q.boost != nil {
		rangeQuery["boost"] = *q.boost
	}

	m := map[string]any{
		"range": map[string]any{
			q.field: rangeQuery,
		},
	}

	return m, nil
}

func (q *TermLevelQueriesRange) MarshalJson() (string, error) {
	return MarshalJson(q)
}
