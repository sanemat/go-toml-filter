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

	tableConfig := tree.Get(tableName).(*toml.Tree)
	tmp := tableConfig.String()
	fmt.Println(tmp)
}
