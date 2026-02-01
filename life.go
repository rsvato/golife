package main

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

const (
	rows              = 10
	cols              = 10
	minNeighborsAlive = 2
)

func round(coord int, around int) int {
	for coord < 0 {
		coord += around
	}
	if coord >= around {
		coord = coord % around
	}
	return coord
}

type Field []bool

func NewField() Field {
	var f = make(Field, cols*rows)
	return f
}

func (f Field) String() string {
	var builder strings.Builder
	for i := range cols {
		for j := range rows {
			if f[i*cols+j] {
				builder.WriteString("█ ")
			} else {
				builder.WriteString("  ")
			}
		}
		builder.WriteString("\n")
	}
	return builder.String()
}

func (f Field) Seed() {
	for i := range f {
		x := rand.IntN(100)
		f[i] = x%4 == 0
	}
}

func (f Field) Alive(x int, y int) bool {
	return f[x*rows+y]
}

func (f Field) AliveNeighbors(x int, y int) int {
	return 0
}

func main() {
	f := NewField()
	f.Seed()
	fmt.Println(f)
}
