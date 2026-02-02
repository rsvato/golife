package main

import (
	"fmt"
	"math/rand/v2"
	"strings"
	"time"
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

type Field struct {
	cols int
	rows int
	data []int
}

func NewField() Field {
	data := make([]int, cols*rows)
	return Field{
		cols: cols,
		rows: rows,
		data: data,
	}
}

func CustomField(cols int, rows int) Field {
	var data = make([]int, cols*rows)
	return Field{
		cols: cols,
		rows: rows,
		data: data,
	}
}

func (f Field) String() string {
	var builder strings.Builder
	for i := range f.cols {
		for j := range f.rows {
			if f.at(i, j) > 0 {
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
	for i := range f.data {
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

func (f Field) setAt(x int, y int, alive bool) {
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
	} else {
		return aliveAround == 3
	}
}

func (f Field) Step() Field {
	newField := Field{
		cols: f.cols,
		rows: f.rows,
		data: make([]int, f.cols*f.rows),
	}
	for i := range f.cols {
		for j := range f.rows {
			newField.setAt(i, j, f.Next(i, j))
		}
	}
	return newField
}

func main() {
	f := CustomField(5, 5)
	f.setAt(2, 1, true)
	f.setAt(2, 2, true)
	f.setAt(2, 3, true)
	fmt.Println(f)
	for _ = range 5 {
		f = f.Step()
		fmt.Println(f)
		time.Sleep(1 * time.Second)
		fmt.Print("\n")
	}

}
