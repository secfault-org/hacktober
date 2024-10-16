package main

import (
	"github.com/secfault-org/hacktober/internal/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
}
