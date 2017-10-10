package jec

import (
	"encoding/json"
	"fmt"
	"regexp"
)

const jgcVarRegBase = `"\$%s"`

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
		re := regexp.MustCompile(fmt.Sprintf(jgcVarRegBase, k))
		b = re.ReplaceAll(b, val)
	}

	return b, nil
}
