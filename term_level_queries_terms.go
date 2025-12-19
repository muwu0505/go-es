package goesdsl

import (
	"errors"
)

/*
	doc: https://www.elastic.co/docs/reference/query-languages/query-dsl/query-dsl-term-query
*/

type TermLevelQueriesTerms struct {
	field string
	value []any
	boost *float64
}

func NewTermsQuery(field string) *TermLevelQueriesTerms {
	return NewTermLevelQueriesTerms(field)
}

func (q *TermLevelQueriesTerms) Valid() error {
	if q == nil {
		return errors.New("nil query")
	}

	if q.field == "" {
		return errors.New("field is required")
	}
	if q.value == nil {
		return errors.New("value is required")
	}

	return nil
}

func (q *TermLevelQueriesTerms) Map() (map[string]any, error) {
	err := q.Valid()
	if err != nil {
		return nil, err
	}

	_map := map[string]any{
		q.field: q.value,
	}

	if q.boost != nil {
		_map["boost"] = *q.boost
	}

	m := map[string]any{
		"terms": _map,
	}

	return m, nil
}

func (q *TermLevelQueriesTerms) Source() (string, error) {
	return Source(q)
}
