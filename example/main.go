package main

import (
	"fmt"
	"log"

	"github.com/ieee0824/jec"
)

func main() {
	base := `
	{
		"Integer": "$int",
		"String": "$string",
		"Object": "$object"
	}
	`

	varsString := `
	{
		"int": 1,
		"string": "hoge",
		"object": {
			"array": ["foo", "bar", "baz"]
		}
	}
	`

	result, err := jec.Embed([]byte(base), []byte(varsString))
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(result))
}
