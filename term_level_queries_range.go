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

func (q *TermLevelQueriesRange) Map() (map[string]any, error) {
	err := q.Valid()
	if err != nil {
		return nil, err
	}

	_map := make(map[string]any)

	if q.gt != nil {
		_map["gt"] = q.gt
	}
	if q.gte != nil {
		_map["gte"] = q.gte
	}
	if q.lt != nil {
		_map["lt"] = q.lt
	}
	if q.lte != nil {
		_map["lte"] = q.lte
	}
	if q.format != "" {
		_map["format"] = q.format
	}
	if q.relation != "" {
		_map["relation"] = q.relation
	}
	if q.timeZone != "" {
		_map["time_zone"] = q.timeZone
	}
	if q.boost != nil {
		_map["boost"] = *q.boost
	}

	m := map[string]any{
		"range": map[string]any{
			q.field: _map,
		},
	}

	return m, nil
}

func (q *TermLevelQueriesRange) Source() (string, error) {
	return Source(q)
}
