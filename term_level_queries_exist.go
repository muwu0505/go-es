package goesdsl

import (
	"errors"
)

/*
	doc: https://www.elastic.co/docs/reference/query-languages/query-dsl/query-dsl-exists-query
*/

type TermLevelQueriesExist struct {
	field string
}

func NewExistQuery(field string) *TermLevelQueriesExist {
	return NewTermLevelQueriesExist(field)
}

func (q *TermLevelQueriesExist) Valid() error {
	if q == nil {
		return errors.New("nil query")
	}

	if q.field == "" {
		return errors.New("field is required")
	}

	return nil
}

func (q *TermLevelQueriesExist) ToMap() (map[string]any, error) {
	err := q.Valid()
	if err != nil {
		return nil, err
	}

	m := map[string]any{
		"exists": map[string]any{
			"field": q.field,
		},
	}

	return m, nil
}

func (q *TermLevelQueriesExist) MarshalJson() (string, error) {
	return MarshalJson(q)
}
