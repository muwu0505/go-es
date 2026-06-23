package goesdsl

import (
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
