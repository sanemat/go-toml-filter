package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pelletier/go-toml"
)

func main() {
	tableName := os.Args[1]

	tree, err := toml.LoadReader(os.Stdin)
	if err != nil {
		log.Fatalf("Error loading TOML: %s", err)
	}

	if !tree.Has(tableName) {
		log.Fatalf("Error tableName does not exist: %s", tableName)
	}

	tableConfig := tree.Get(tableName).(*toml.Tree)
	m := make(map[string]interface{})
	for _, element := range tableConfig.Keys() {
		m[element] = tableConfig.Get(element)
	}
	tmp := convert(m)
	fmt.Println(tmp)
}

func convert(m map[string]interface{}) string {
	return fmt.Sprint(m)
}
