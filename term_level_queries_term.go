package goesdsl

import (
	"errors"
)

/*
	doc: https://www.elastic.co/docs/reference/query-languages/query-dsl/query-dsl-term-query
*/

type TermLevelQueriesTerm struct {
	field           string
	value           any
	boost           *float64
	caseInsensitive *bool
}

func NewTermQuery(field string) *TermLevelQueriesTerm {
	return NewTermLevelQueriesTerm(field)
}

func (q *TermLevelQueriesTerm) Valid() error {
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

func (q *TermLevelQueriesTerm) Map() (map[string]any, error) {
	err := q.Valid()
	if err != nil {
		return nil, err
	}

	term := map[string]any{
		"value": q.value,
	}

	if q.boost != nil {
		term["boost"] = *q.boost
	}
	if q.caseInsensitive != nil {
		term["case_insensitive"] = *q.caseInsensitive
	}

	m := map[string]any{
		"term": map[string]any{
			q.field: term,
		},
	}

	return m, nil
}

func (q *TermLevelQueriesTerm) Source() (string, error) {
	return Source(q)
}
