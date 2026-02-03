package main

import (
	"fmt"
	"math/rand/v2"
	"strings"
	"time"
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

func ReadStrings(s string) *Field {
	rows := strings.Split(strings.TrimSpace(s), "\n")
	var cleanRows []string

    for _, r := range rows {
        trimmed := strings.TrimSpace(r)
        if trimmed != "" {
            cleanRows = append(cleanRows, trimmed)
        }
    }

    if len(cleanRows) == 0 {
    	return &Field{
     		cols: 0,
       		rows: 0,
         	data: make([]int, 0),
     	}
    }

	height := len(cleanRows)
	width := len(cleanRows[0])
	data := make([]int, height * width)
	result := Field {
		cols: width,
		rows: height,
		data: data,
	}
	for i, row := range cleanRows {
		row = strings.TrimRight(row, "\r")
		for j, ch := range row {
			alive := false
			if (ch == '#') {
				alive = true
			}
			result.setAt(i, j, alive)
		}
	}
	return &result
}

func main() {
	f := NewField(5, 7)
	f.setAt(2, 1, true)
	f.setAt(2, 2, true)
	f.setAt(2, 3, true)
	fmt.Println(f)
	for i := 0; i < 5; i++ {
		fmt.Print("\033[H\033[2J")
		f = f.Step()
		fmt.Println(f)
		time.Sleep(1 * time.Second)
	}

}
