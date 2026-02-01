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

type Field []byte

func NewField() Field {
	var f = make(Field, cols*rows)
	return f
}

func (f Field) String() string {
	var builder strings.Builder
	for i := range cols {
		for j := range rows {
			if f[i*rows+j] > 0 {
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
		if x%4 == 0 {
			f[i] = 1
		} else {
			f[i] = 0
		}
	}
}

func (f Field) Alive(x int, y int) bool {
	return f[idx(x, y)] == 1
}

func idx(x int, y int) int {
	return y*rows + x
}

func (f Field) AliveNeighbors(x int, y int) int {
	abs := idx(x, y)
	start := abs - 4
	end := abs + 4
	var result byte
	for start <= end {
		if start == abs {
			start++
			continue
		}
		var idx int
		if start < 0 {
			idx = start + rows*cols
		} else if start >= rows*cols {
			idx = start % (rows * cols)
		} else {
			idx = start
		}
		result += f[idx]
		start++
	}
	return int(result)
}

func main() {
	f := NewField()
	f.Seed()
	fmt.Println(f)
}
