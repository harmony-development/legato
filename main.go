package main

import (
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	if err := Start(); err != nil {
		fmt.Println(err)
	}
}

// Start starts the server.
func Start() error {
	_ = godotenv.Load()

	cfg, err := ReadConfig("./config.yml")
	if err != nil {
		return err
	}

	return nil
}
