package main

import (
	"fmt"
	"log"
	"os"

	"strings"

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
	tmp := tableConfig.ToMap()
	var b strings.Builder
	for key, value := range tmp {
		if b.Len() != 0 {
			b.WriteString("\n")
		}
		fmt.Fprintf(&b, "%v=%v", key, value)
	}
	fmt.Println(b.String())
}
