package goesdsl

import (
	"errors"
)

/*
	doc: https://www.elastic.co/docs/reference/query-languages/query-dsl/query-dsl-fuzzy-query
*/

type TermLevelQueriesFuzzy struct {
	field          string
	value          string
	fuzziness      string
	maxExpansions  *int32
	prefixLength   *int32
	transpositions *bool
	rewrite        string
}

func NewFuzzyQuery(field string) *TermLevelQueriesFuzzy {
	return NewTermLevelQueriesFuzzy(field)
}

func (q *TermLevelQueriesFuzzy) Valid() error {
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

func (q *TermLevelQueriesFuzzy) Map() (map[string]any, error) {
	err := q.Valid()
	if err != nil {
		return nil, err
	}

	fuzzy := map[string]any{
		"value": q.value,
	}

	if q.fuzziness != "" {
		fuzzy["fuzziness"] = q.fuzziness
	}
	if q.maxExpansions != nil {
		fuzzy["max_expansions"] = *q.maxExpansions
	}
	if q.prefixLength != nil {
		fuzzy["prefix_length"] = *q.prefixLength
	}
	if q.transpositions != nil {
		fuzzy["transpositions"] = *q.transpositions
	}
	if q.rewrite != "" {
		fuzzy["rewrite"] = q.rewrite
	}

	m := map[string]any{
		"fuzzy": map[string]any{
			q.field: fuzzy,
		},
	}

	return m, nil
}

func (q *TermLevelQueriesFuzzy) Source() (string, error) {
	return Source(q)
}
