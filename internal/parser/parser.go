package parser

import (
	"io"
)

type Parser interface {
	Parse(reader io.Reader)
}
