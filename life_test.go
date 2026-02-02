package main

import (
	"testing"
)

func TestAliveNeighbors(t *testing.T) {
	f := NewField()
	f.data[0] = 1
	x := f.AliveNeighbors(0, 0)
	if x != 0 {
		t.Errorf("On one cell desk all cells around must not be alive")
	}
	x = f.AliveNeighbors(9, 9)
	if x != 1 {
		t.Errorf("At least one cell must be alive")
	}
}

func TestAliveNeighbors2(t *testing.T) {
	f := CustomField(5, 5)
	f.setAt(2, 1, true)
	f.setAt(2, 2, true)
	f.setAt(2, 3, true)

	x := f.AliveNeighbors(1, 3)
	if x != 2 {
		t.Errorf("Alive neighbors must be 2 at point 1,3")
	}
}

func TestAt(t *testing.T) {
	f := CustomField(3, 3)
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
	f := CustomField(5, 5)
	f.setAt(2, 1, true)
	f.setAt(2, 2, true)
	f.setAt(2, 3, true)
	alive := f.Next(2, 1)
	if alive {
		t.Errorf("First cell in line should be dead at next move, alive nbs %d", f.AliveNeighbors(2, 1))
	}
	alive = f.Next(2, 2)
	if !alive {
		t.Errorf("Central cell has two neighbors, so it should be alive on next move")
	}

	alive = f.Next(2, 3)

	if alive {
		t.Errorf("Last cell has one neighbor, so it should be dead on next move, %d", f.AliveNeighbors(2, 3))
	}

	alive = f.Next(1, 2)
	if !alive {
		t.Errorf("Cell above center must be alive one, %d", f.AliveNeighbors(1, 2))
	}
	alive = f.Next(3, 2)
	if !alive {
		t.Errorf("Cell below center must be alive too, %d", f.AliveNeighbors(3, 2))
	}
}
