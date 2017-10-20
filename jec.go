package jec

import (
	"encoding/json"
	"fmt"
	"strings"
)

const jgcBase = `"$%s"`

func prettify(b []byte) ([]byte, error) {
	var i interface{}
	if err := json.Unmarshal(b, &i); err != nil {
		return nil, err
	}
	return json.MarshalIndent(i, "", "  ")
}

func Embed(b, v []byte) ([]byte, error) {
	vars := map[string]interface{}{}

	if err := json.Unmarshal(v, &vars); err != nil {
		return nil, err
	}

	for k, v := range vars {
		val, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}

		b = []byte(strings.Replace(string(b), fmt.Sprintf(jgcBase, k), string(val), -1))
	}

	return prettify(b)
}
