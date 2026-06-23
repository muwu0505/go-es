package goesdsl

import (
	"errors"
)

/*
	doc: https://www.elastic.co/docs/reference/query-languages/query-dsl/query-dsl-regexp-query
*/

type TermLevelQueriesRegexp struct {
	field                 string
	value                 any
	flags                 string
	caseInsensitive       *bool
	maxDeterminizedStates *int64
	rewrite               string
}

func NewRegexpQuery(field string) *TermLevelQueriesRegexp {
	return NewTermLevelQueriesRegexp(field)
}

func (q *TermLevelQueriesRegexp) Valid() error {
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

func (q *TermLevelQueriesRegexp) ToMap() (map[string]any, error) {
	err := q.Valid()
	if err != nil {
		return nil, err
	}

	regexp := map[string]any{
		"value": q.value,
	}

	if q.flags != "" {
		regexp["flags"] = q.flags
	}
	if q.caseInsensitive != nil {
		regexp["case_insensitive"] = *q.caseInsensitive
	}
	if q.maxDeterminizedStates != nil {
		regexp["max_determinized_states"] = *q.maxDeterminizedStates
	}
	if q.rewrite != "" {
		regexp["rewrite"] = q.rewrite
	}

	m := map[string]any{
		"regexp": map[string]any{
			q.field: regexp,
		},
	}

	return m, nil
}

func (q *TermLevelQueriesRegexp) MarshalJson() (string, error) {
	return MarshalJson(q)
}
