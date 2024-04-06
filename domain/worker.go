package domain

import (
	"bufio"
	"io"
)

type Worker struct {
	ID           int
	Stdin        io.WriteCloser
	StdOutReader *bufio.Reader
}