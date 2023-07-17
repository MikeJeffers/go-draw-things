package main

import (
	"fmt"
	"os"
	"testing"
)

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

func TestNorm(t *testing.T) {
	got := norm(1, 10)
	if got != 1 {
		t.Fail()
	}

	got = norm(11, 10)
	if got != 1 {
		t.Fail()
	}

	got = norm(-1, 10)
	if got != 9 {
		t.Fail()
	}

	got = norm(-11, 10)
	if got != 9 {
		t.Fail()
	}

	for i := -10; i < 10; i++ {
		got = norm(i, 5)
		if got < 0 || got > 5 {
			t.Fail()
		}
		fmt.Println(i, 5, got)
	}
}

func TestMain(m *testing.M) {

	os.Exit(m.Run())
}
