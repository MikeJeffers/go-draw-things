package main

import "testing"

func TestCrossing(t *testing.T) {
	path := Path{points: []Point[int]{{x: 1, y: 1}, {x: 2, y: 2}}}
	got := checkDiagonalCross(path, 1, 2, 2, 1)
	if !got {
		t.Error()
	}
	got = checkDiagonalCross(path, 2, 1, 1, 2)
	if !got {
		t.Error()
	}

	got = checkDiagonalCross(path, 3, 4, 5, 5)
	if got {
		t.Error()
	}

	got = checkDiagonalCross(path, 1, 1, 2, 2)
	if got {
		t.Error()
	}
}

func TestInitGrid(t *testing.T) {
	g := initGrid(10, 10)
	if len(g) != 10 || len(g[0]) != 10 {
		t.Errorf("Unexpected matrix dimensions")
	}
}
