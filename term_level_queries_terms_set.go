package goesdsl

import (
	"errors"
)

/*
	doc: https://www.elastic.co/docs/reference/query-languages/query-dsl/query-dsl-terms-query
*/

type TermLevelQueriesTermsSet struct {
	field                    string
	terms                    []any
	minimumShouldMatch       *float64
	minimumShouldMatchField  string
	minimumShouldMatchScript *Script
}

func NewTermsSetQuery(field string) *TermLevelQueriesTermsSet {
	return NewTermLevelQueriesTermsSet(field)
}

func (q *TermLevelQueriesTermsSet) Valid() error {
	if q == nil {
		return errors.New("nil query")
	}

	if q.field == "" {
		return errors.New("field is required")
	}
	if q.terms == nil {
		return errors.New("terms is required")
	}

	return nil
}

func (q *TermLevelQueriesTermsSet) Map() (map[string]any, error) {
	err := q.Valid()
	if err != nil {
		return nil, err
	}

	_map := map[string]any{
		"terms": q.terms,
	}

	if q.minimumShouldMatch != nil {
		_map["minimum_should_match"] = *q.minimumShouldMatch
	}
	if q.minimumShouldMatchField != "" {
		_map["minimum_should_match_field"] = q.minimumShouldMatchField
	}
	if q.minimumShouldMatchScript != nil {
		_map["minimum_should_match_script"], err = q.minimumShouldMatchScript.Map()
		if err != nil {
			return nil, err
		}
	}

	m := map[string]any{
		"terms_set": map[string]any{
			q.field: _map,
		},
	}

	return m, nil
}

func (q *TermLevelQueriesTermsSet) Source() (string, error) {
	return Source(q)
}
