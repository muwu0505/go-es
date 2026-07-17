package goesdsl

import (
	"fmt"

	"github.com/bytedance/sonic"
)

type IQuery interface {
	Valid() error
	ToMap() (map[string]any, error)
	MarshalJson() (string, error)
}

func MarshalJson(q IQuery) (string, error) {
	_map, err := q.ToMap()
	if err != nil {
		return "", err
	}

	return sonic.MarshalString(_map)
}

func ValidQueries(name string, queries ...IQuery) error {
	for i, query := range queries {
		if query == nil {
			return fmt.Errorf("%s[%d] query is required", name, i)
		}
		if err := query.Valid(); err != nil {
			return fmt.Errorf("%s[%d]: %w", name, i, err)
		}
	}

	return nil
}

func QueriesToMap(name string, queries ...IQuery) ([]map[string]any, error) {
	if len(queries) == 0 {
		return nil, nil
	}

	qs := make([]map[string]any, 0, len(queries))
	for i, query := range queries {
		_query, err := query.ToMap()
		if err != nil {
			return nil, fmt.Errorf("%s[%d]: %w", name, i, err)
		}
		qs = append(qs, _query)
	}

	return qs, nil
}

type QueryName struct {
	name string
}
