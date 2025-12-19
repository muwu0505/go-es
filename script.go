package goesdsl

import (
	"errors"
)

/*
	doc: https://www.elastic.co/docs/explore-analyze/scripting/modules-scripting-painless
*/

type Script struct {
	lang       string         // Specifies the language the script is written in. Defaults to painless.
	source, id string         // The script itself, which you specify as source for an inline script or id for a stored script.
	params     map[string]any // Specifies any named parameters that are passed into the script as variables.
}

func NewPainlessScript(source, id string, params map[string]any) *Script {
	return &Script{
		source: source,
		id:     id,
		params: params,
	}
}

func NewPainlessSourceScript(source string, params map[string]any) *Script {
	return NewPainlessScript(source, "", params)
}

func NewPainlessIDScript(id string, params map[string]any) *Script {
	return NewPainlessScript("", id, params)
}

func (q *Script) Valid() error {
	if q == nil {
		return errors.New("nil query")
	}

	return nil
}

func (q *Script) Map() (map[string]any, error) {
	err := q.Valid()
	if err != nil {
		return nil, err
	}

	_map := make(map[string]any)

	if q.lang != "" {
		_map["lang"] = q.lang
	}
	if q.source != "" {
		_map["source"] = q.source
	}
	if q.id != "" {
		_map["id"] = q.id
	}
	if len(q.params) > 0 {
		_map["params"] = q.params
	}

	m := map[string]any{
		"script": _map,
	}

	return m, nil
}

func (q *Script) Source() (string, error) {
	return Source(q)
}
