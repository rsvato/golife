package main

import (
	"fmt"
	"strings"
)

const (
	width  = 80
	height = 40
)

type Field [][]bool

func newField() Field {
	var f = make(Field, height)
	for i := range f {
		f[i] = make([]bool, width)
	}
	return f
}

func (f Field) Show() string {
	var builder strings.Builder
	for i := range f {
		for j := range f[i] {
			if f[i][j] {
				builder.WriteString("x")
			} else {
				builder.WriteString(".")
			}
		}
		builder.WriteString("\n")
	}
	return builder.String()
}

func main() {
	f := newField()
	fmt.Println(f.Show())
}
