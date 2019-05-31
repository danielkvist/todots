package main

import (
	"log"

	"github.com/danielkvist/todots/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatalf("%v", err)
	}
}
