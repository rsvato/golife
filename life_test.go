package main

import (
	"testing"
)

func TestAliveNeighbors(t *testing.T) {
	f := NewField()
	f[0] = 1
	x := f.AliveNeighbors(0, 0)
	if x != 0 {
		t.Errorf("On one cell desk all cells around must not be alive")
	}
	x = f.AliveNeighbors(9, 9)
	if x != 1 {
		t.Errorf("At least one cell must be alive")
	}
}
