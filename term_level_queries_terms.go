package goesdsl

import (
	"errors"
)

/*
	doc: https://www.elastic.co/docs/reference/query-languages/query-dsl/query-dsl-terms-query
*/

type TermLevelQueriesTerms struct {
	field  string
	value  []any
	lookup *TermsLookup
	boost  *float64
}

type TermsLookup struct {
	Index   string
	ID      string
	Path    string
	Routing string
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
	if len(q.value) == 0 && q.lookup == nil {
		return errors.New("value is required")
	}
	if q.lookup != nil {
		if q.lookup.Index == "" {
			return errors.New("lookup index is required")
		}
		if q.lookup.ID == "" {
			return errors.New("lookup id is required")
		}
		if q.lookup.Path == "" {
			return errors.New("lookup path is required")
		}
	}

	return nil
}

func (q *TermLevelQueriesTerms) ToMap() (map[string]any, error) {
	err := q.Valid()
	if err != nil {
		return nil, err
	}

	terms := map[string]any{}

	if q.lookup != nil {
		lookup := map[string]any{
			"index": q.lookup.Index,
			"id":    q.lookup.ID,
			"path":  q.lookup.Path,
		}
		if q.lookup.Routing != "" {
			lookup["routing"] = q.lookup.Routing
		}
		terms[q.field] = lookup
	} else {
		terms[q.field] = q.value
	}
	if q.boost != nil {
		terms["boost"] = *q.boost
	}

	m := map[string]any{
		"terms": terms,
	}

	return m, nil
}

func (q *TermLevelQueriesTerms) MarshalJson() (string, error) {
	return MarshalJson(q)
}
