package main

import (
	"os"

	"github.com/fedeizzo/gonetworth/internal/parser"
)

func main() {
	p := parser.NewN26Parser()
	f, err := os.Open("./transactions.csv")
	if err != nil {
		panic(err)
	}

	p.Parse(f)
}
