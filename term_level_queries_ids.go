package goesdsl

import (
	"errors"
)

/*
	doc: https://www.elastic.co/docs/reference/query-languages/query-dsl/query-dsl-ids-query
*/

type TermLevelQueriesIDs struct {
	values []string
}

func NewIDsQuery() *TermLevelQueriesIDs {
	return NewTermLevelQueriesIDs()
}

func (q *TermLevelQueriesIDs) Valid() error {
	if q == nil {
		return errors.New("nil query")
	}

	if len(q.values) <= 0 {
		return errors.New("values is required")
	}

	return nil
}

func (q *TermLevelQueriesIDs) ToMap() (map[string]any, error) {
	err := q.Valid()
	if err != nil {
		return nil, err
	}

	return map[string]any{
		"ids": map[string]any{
			"values": q.values,
		},
	}, nil
}

func (q *TermLevelQueriesIDs) MarshalJson() (string, error) {
	return MarshalJson(q)
}
