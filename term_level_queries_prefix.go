package goesdsl

import (
	"errors"
)

/*
	doc: https://www.elastic.co/docs/reference/query-languages/query-dsl/query-dsl-prefix-query
*/

type TermLevelQueriesPrefix struct {
	field           string
	value           string
	rewrite         string
	caseInsensitive *bool
}

func NewPrefixQuery(field string) *TermLevelQueriesPrefix {
	return NewTermLevelQueriesPrefix(field)
}

func (q *TermLevelQueriesPrefix) Valid() error {
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

func (q *TermLevelQueriesPrefix) ToMap() (map[string]any, error) {
	err := q.Valid()
	if err != nil {
		return nil, err
	}

	prefix := map[string]any{
		"value": q.value,
	}

	if q.rewrite != "" {
		prefix["rewrite"] = q.rewrite
	}
	if q.caseInsensitive != nil {
		prefix["case_insensitive"] = *q.caseInsensitive
	}

	return map[string]any{
		"prefix": map[string]any{
			q.field: prefix,
		},
	}, nil
}

func (q *TermLevelQueriesPrefix) MarshalJson() (string, error) {
	return MarshalJson(q)
}
