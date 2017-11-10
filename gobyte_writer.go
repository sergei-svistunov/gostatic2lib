package main

import (
	"fmt"
	"io"
)

type GoByteWriter struct {
	w io.Writer
}

func NewGoByteWriter(w io.Writer) *GoByteWriter {
	return &GoByteWriter{w}
}

func (w *GoByteWriter) Write(p []byte) (n int, err error) {
	for i, b := range p {
		if i > 0 && i%30 == 0 {
			_, err := w.w.Write([]byte("\n"))
			if err != nil {
				return i, err
			}
		}

		w.w.Write([]byte(fmt.Sprintf("%d, ", b)))
		if err != nil {
			return i, err
		}
	}

	return len(p), nil
}
