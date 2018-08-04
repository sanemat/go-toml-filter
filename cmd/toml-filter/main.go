package main

import (
	"fmt"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

func main() {
	var tmp interface{}
	if _, err := toml.DecodeReader(os.Stdin, &tmp); err != nil {
		log.Fatalf("Error decoding TOML: %s", err)
	}
	fmt.Println("hello world")
}
