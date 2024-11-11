package parser

import (
	"fmt"
	"io"

	"github.com/gocarina/gocsv"
)

type transaction struct {
	Date   string  `csv:"Value Date"`
	Actor  string  `csv:"Partner Name"`
	Type   string  `csv:"Type"`
	Amount float64 `csv:"Amount (EUR)"`
}

type n26 struct{}

func NewN26Parser() n26 {
	return n26{}
}

func (p n26) Parse(reader io.Reader) {
	var transactions []transaction

	if err := gocsv.Unmarshal(reader, &transactions); err != nil {
		panic(err)
	}

	for _, t := range transactions {
		fmt.Println(t)
	}
}
