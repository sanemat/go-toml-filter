package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pelletier/go-toml"
)

func main() {
	fmt.Println(os.Args[1])

	tree, err := toml.LoadReader(os.Stdin)
	if err != nil {
		log.Fatalf("Error loading TOML: %s", err)
	}
	tmp := tree.String()
	fmt.Println(tmp)
}
