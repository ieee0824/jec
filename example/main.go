package main

import (
	"fmt"
	"log"

	"github.com/ieee0824/jgc"
)

func main() {
	base := `
	{
		"Integer": "$int",
		"String": "$string"
	}
	`

	varsString := `
	{
		"int": 1,
		"string": "hoge"
	}
	`

	result, err := jgc.Embedde([]byte(base), []byte(varsString))
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(result))
}
