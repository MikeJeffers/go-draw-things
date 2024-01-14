package main

import "testing"

func TestCrossing(t *testing.T) {
	path := Path{points: []Point[int]{{x: 1, y: 1}, {x: 2, y: 2}}}
	got := checkDiagonalCross(path, 1, 2, 2, 1)
	if !got {
		t.Fail()
	}
	got = checkDiagonalCross(path, 2, 1, 1, 2)
	if !got {
		t.Fail()
	}

	got = checkDiagonalCross(path, 3, 4, 5, 5)
	if got {
		t.Fail()
	}

	got = checkDiagonalCross(path, 1, 1, 2, 2)
	if got {
		t.Fail()
	}
}
