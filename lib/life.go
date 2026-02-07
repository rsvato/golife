package lib

import (
	"math/rand/v2"
	"strings"
)

type Field struct {
	cols int
	rows int
	data []int
}

func NewField(cols int, rows int) Field {
	data := make([]int, cols*rows)
	return Field{
		cols: cols,
		rows: rows,
		data: data,
	}
}

func (f Field) String() string {
	var builder strings.Builder
	builder.Grow(f.rows * (f.cols*2 + 1))
	for i := 0; i < f.rows; i++ {
		for j := 0; j < f.cols; j++ {
			if f.Alive(i, j) {
				builder.WriteString("█ ")
			} else {
				builder.WriteString("  ")
			}
		}
		builder.WriteString("\n")
	}
	return builder.String()
}

func (f *Field) Seed() {
	for i := 0; i < len(f.data); i++ {
		x := rand.IntN(100)
		if x%4 == 0 {
			f.data[i] = 1
		} else {
			f.data[i] = 0
		}
	}
}

func (f Field) Alive(x int, y int) bool {
	return f.at(x, y) == 1
}

func (f Field) at(x int, y int) int {
	idx := x*f.cols + y
	return f.data[idx]
}

func (f *Field) setAt(x int, y int, alive bool) {
	idx := x*f.cols + y
	if alive {
		f.data[idx] = 1
	} else {
		f.data[idx] = 0
	}
}

func (f Field) AliveNeighbors(x int, y int) int {
	var result int
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}

			nx := (x + dx + f.rows) % f.rows
			ny := (y + dy + f.cols) % f.cols

			result += f.at(nx, ny)
		}
	}
	return result
}

func (f Field) Next(x int, y int) bool {
	aliveAround := f.AliveNeighbors(x, y)
	alive := f.Alive(x, y)
	if alive {
		return aliveAround == 2 || aliveAround == 3
	}
	return aliveAround == 3
}

func (f Field) Step() Field {
	newField := Field{
		cols: f.cols,
		rows: f.rows,
		data: make([]int, f.cols*f.rows),
	}
	for i := 0; i < f.rows; i++ {
		for j := 0; j < f.cols; j++ {
			newField.setAt(i, j, f.Next(i, j))
		}
	}
	return newField
}
