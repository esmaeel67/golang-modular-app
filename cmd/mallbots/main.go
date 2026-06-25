package main

import (
	"fmt"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Println("startUp", err.Error())
		os.Exit(1)
	}
}

func run() error {

	return nil
}
