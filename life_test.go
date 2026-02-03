package main

import (
	"testing"
)

func TestAliveNeighbors(t *testing.T) {
	tests := []struct {
		name   string
		alive  [][2]int
		x, y   int
		expect int
	}{
		{
			name: "single neighbor",
			alive: [][2]int{
				{1, 1},
			},
			x:      0,
			y:      0,
			expect: 1,
		},
		{
			name: "three neighbors",
			alive: [][2]int{
				{0, 1},
				{1, 0},
				{1, 1},
			},
			x:      0,
			y:      0,
			expect: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewField(5, 5)
			for _, c := range tt.alive {
				f.setAt(c[0], c[1], true)
			}
			got := f.AliveNeighbors(tt.x, tt.y)
			if got != tt.expect {
				t.Fatalf("expected %d, got %d", tt.expect, got)
			}
		})
	}
}

func TestAt(t *testing.T) {
	f := NewField(3, 3)
	f.data[0] = 1
	f.data[8] = 2
	x := f.at(0, 0)
	if x != 1 {
		t.Errorf("First cell must be 1, but is %d", x)
	}
	y := f.at(2, 2)
	if y != 2 {
		t.Errorf("Last cell must be 2, but is %d", y)
	}
}

func TestNext(t *testing.T) {
	tests := []struct {
		name   string
		alive  bool
		around int
		expect bool
	}{
		{"alive with 2 neighbors survives", true, 2, true},
		{"alive with 3 neighbors survives", true, 3, true},
		{"alive with 1 neighbor dies", true, 1, false},
		{"dead with 3 neighbors becomes alive", false, 3, true},
		{"dead with 2 neighbors stays dead", false, 2, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewField(3, 3)
			if tt.alive {
				f.setAt(1, 1, true)
			}

			count := 0
			for i := 0; count < tt.around; i++ {
				f.setAt(0, i, true)
				count++
			}

			got := f.Next(1, 1)
			if got != tt.expect {
				t.Fatalf("expected %v, got %v", tt.expect, got)
			}
		})
	}
}

func TestStep_Blinker(t *testing.T) {
	f := NewField(5, 5)

	f.setAt(2, 1, true)
	f.setAt(2, 2, true)
	f.setAt(2, 3, true)

	f = f.Step()

	expected := [][2]int{
		{1, 2},
		{2, 2},
		{3, 2},
	}

	for _, c := range expected {
		if !f.Alive(c[0], c[1]) {
			t.Fatalf("expected cell %v to be alive", c)
		}
	}
}

func TestStep_Block(t *testing.T) {
	f := NewField(4, 4)

	block := [][2]int{
		{1, 1},
		{1, 2},
		{2, 1},
		{2, 2},
	}

	for _, c := range block {
		f.setAt(c[0], c[1], true)
	}

	f2 := f.Step()

	for _, c := range block {
		if !f2.Alive(c[0], c[1]) {
			t.Fatalf("block cell %v should stay alive", c)
		}
	}
}

func TestReadStrings(t *testing.T) {
	input := `
	.#.
	.#.
	.#.
	`
	f := ReadStrings(input)
	if f.cols != 3 {
		t.Fatalf("Cols should be 3, but is %d", f.cols)
	}
	if f.rows != 3 {
		t.Fatalf("Rows should be 3, but is %d", f.rows)
	}
}
