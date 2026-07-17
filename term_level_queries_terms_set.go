package goesdsl

import (
	"errors"
)

/*
	doc: https://www.elastic.co/docs/reference/query-languages/query-dsl/query-dsl-terms-set-query
*/

type TermLevelQueriesTermsSet struct {
	field                    string
	terms                    []any
	minimumShouldMatch       any
	minimumShouldMatchField  string
	minimumShouldMatchScript *Script
	boost                    *float64
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
	if len(q.terms) == 0 {
		return errors.New("terms is required")
	}
	minimumShouldMatchOptions := 0
	if q.minimumShouldMatch != nil {
		minimumShouldMatchOptions++
	}
	if q.minimumShouldMatchField != "" {
		minimumShouldMatchOptions++
	}
	if q.minimumShouldMatchScript != nil {
		minimumShouldMatchOptions++
	}
	if minimumShouldMatchOptions == 0 {
		return errors.New("minimum_should_match is required")
	}
	if minimumShouldMatchOptions > 1 {
		return errors.New("only one minimum_should_match option is allowed")
	}

	return nil
}

func (q *TermLevelQueriesTermsSet) ToMap() (map[string]any, error) {
	err := q.Valid()
	if err != nil {
		return nil, err
	}

	termsSet := map[string]any{
		"terms": q.terms,
	}

	if q.minimumShouldMatch != nil {
		termsSet["minimum_should_match"] = q.minimumShouldMatch
	}
	if q.minimumShouldMatchField != "" {
		termsSet["minimum_should_match_field"] = q.minimumShouldMatchField
	}
	if q.minimumShouldMatchScript != nil {
		termsSet["minimum_should_match_script"], err = q.minimumShouldMatchScript.ToMap()
		if err != nil {
			return nil, err
		}
	}
	if q.boost != nil {
		termsSet["boost"] = *q.boost
	}

	return map[string]any{
		"terms_set": map[string]any{
			q.field: termsSet,
		},
	}, nil
}

func (q *TermLevelQueriesTermsSet) MarshalJson() (string, error) {
	return MarshalJson(q)
}
