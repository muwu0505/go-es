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
	wildcard        string
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
	if q.value == "" && q.wildcard == "" {
		return errors.New("value is required")
	}

	return nil
}

func (q *TermLevelQueriesWildcard) Map() (map[string]any, error) {
	err := q.Valid()
	if err != nil {
		return nil, err
	}

	wildcard := map[string]any{}
	if q.wildcard != "" {
		wildcard["wildcard"] = q.wildcard
	} else {
		wildcard["value"] = q.value
	}

	if q.boost != nil {
		wildcard["boost"] = *q.boost
	}
	if q.caseInsensitive != nil {
		wildcard["case_insensitive"] = *q.caseInsensitive
	}
	if q.rewrite != "" {
		wildcard["rewrite"] = q.rewrite
	}

	m := map[string]any{
		"wildcard": map[string]any{
			q.field: wildcard,
		},
	}

	return m, nil
}

func (q *TermLevelQueriesWildcard) Source() (string, error) {
	return Source(q)
}
