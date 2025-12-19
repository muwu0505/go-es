package goesdsl

import (
	"errors"
)

/*
	doc: https://www.elastic.co/docs/reference/query-languages/query-dsl/query-dsl-wildcard-query
*/

type TermLevelQueriesWildcard struct {
	field           string
	value           string
	boost           *float64
	caseInsensitive *bool
	rewrite         string
}

func NewWildcardQuery(field string) *TermLevelQueriesWildcard {
	return NewTermLevelQueriesWildcard(field)
}

func (q *TermLevelQueriesWildcard) Valid() error {
	if q == nil {
		return errors.New("nil query")
	}

	if q.field == "" {
		return errors.New("field is required")
	}
	if q.value == "" {
		return errors.New("value is required")
	}

	return nil
}

func (q *TermLevelQueriesWildcard) Map() (map[string]any, error) {
	err := q.Valid()
	if err != nil {
		return nil, err
	}

	_map := map[string]any{
		"value": q.value,
	}

	if q.boost != nil {
		_map["boost"] = *q.boost
	}
	if q.caseInsensitive != nil {
		_map["case_insensitive"] = *q.caseInsensitive
	}
	if q.rewrite != "" {
		_map["rewrite"] = q.rewrite
	}

	m := map[string]any{
		"wildcard": map[string]any{
			q.field: _map,
		},
	}

	return m, nil
}

func (q *TermLevelQueriesWildcard) Source() (string, error) {
	return Source(q)
}
