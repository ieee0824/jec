package jec

import (
	"encoding/json"
	"reflect"
	"testing"
)

func compaireJSON(a, b []byte) bool {
	var am, bm = map[string]interface{}{}, map[string]interface{}{}

	if err := json.Unmarshal(a, &am); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(b, &bm); err != nil {
		panic(err)
	}

	return reflect.DeepEqual(am, bm) || (len(am) == 0 && len(bm) == 0)
}

func TestEmbedde(t *testing.T) {
	tests := []struct {
		base string
		val  string
		want string
		err  bool
	}{
		{
			`{}`,
			`{}`,
			`{}`,
			false,
		},
		{
			`
			{
				"Integer": "$int"
			}
			`,
			`
			{
				"int": 1
			}
			`,
			`
			{
				"Integer": 1
			}
			`,
			false,
		},
		{
			`
			{
				"String": "$str"
			}
			`,
			`
			{
				"str": "hoge"
			}
			`,
			`
			{
				"String": "hoge"
			}
			`,
			false,
		},
		{
			`
			{
				"Bool": "$bool"
			}
			`,
			`
			{
				"bool": true
			}
			`,
			`
			{
				"Bool": true
			}
			`,
			false,
		},
	}

	for _, test := range tests {
		got, err := Embedde([]byte(test.base), []byte(test.val))
		if !test.err && err != nil {
			t.Fatalf("should not be error for %v, %v but: %v", test.base, test.val, err)
		}
		if test.err && err == nil {
			t.Fatalf("should be error for %v, %v but not:", test.base, test.val)
		}
		if !compaireJSON(got, []byte(test.want)) {
			t.Fatalf("want %q, but %q:", test.want, got)
		}
	}
}
