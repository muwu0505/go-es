package goesdsl

import (
	"github.com/bytedance/sonic"
)

type IQuery interface {
	Valid() error
	Map() (map[string]any, error)
	Source() (string, error)
}

func Source(q IQuery) (string, error) {
	_map, err := q.Map()
	if err != nil {
		return "", err
	}

	return sonic.MarshalString(_map)
}
