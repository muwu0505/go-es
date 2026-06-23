package goesdsl

import (
	"fmt"
	"testing"
)

func TestQueryExist_Valid(t *testing.T) {
	tests := []*TestBase{
		//{
		//	name:  "valid query",
		//	query: NewExistQuery("user.id"),
		//	err:   false,
		//},
		//{
		//	name:  "missing field",
		//	query: NewExistQuery(""),
		//	err:   true,
		//	want:  "field is required",
		//},
		{
			name:  "basic query",
			query: NewExistQuery("tags"),
			want: map[string]any{
				"exists": map[string]any{
					"field": "tags",
				},
			},
			err: false,
		},
	}

	//TestValid(t, tests)

	for _, test := range tests {
		test.query.MarshalJson()
	}

	TestMap(t, tests)
	fmt.Println(1)
}

func TestQueryExist_Map(t *testing.T) {
	tests := []*TestBase{
		{
			name:  "basic query",
			query: NewExistQuery("tags"),
			want: map[string]any{
				"exists": map[string]any{
					"field": "tags",
				},
			},
			err: false,
		},
	}

	TestMap(t, tests)
}
