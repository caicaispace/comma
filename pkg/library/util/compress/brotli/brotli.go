package brotli

import (
	"bytes"
	"fmt"

	"github.com/andybalholm/brotli"
)

// Test basic encoder usage.
func New() {
	input := []byte("<html><body><H1>Hello world</H1></body></html>")
	out := bytes.Buffer{}
	e := brotli.NewWriterOptions(&out, brotli.WriterOptions{Quality: 1})
	n, err := e.Write(input)
	if err != nil {
	}
	fmt.Println("n", n)
}
