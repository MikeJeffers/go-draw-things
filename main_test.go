package main

import (
	"fmt"
	"os"
	"testing"
)

func TestCrossing(t *testing.T) {
	path := Path{points: []Point[int]{{x: 1, y: 1}, {x: 2, y: 2}}}
	got := checkDiagonalCross(path, 1, 2, 2, 1)
	fmt.Println(got)
	if !got {
		t.Fail()
	}
	got = checkDiagonalCross(path, 2, 1, 1, 2)
	fmt.Println(got)
	if !got {
		t.Fail()
	}

	got = checkDiagonalCross(path, 3, 4, 5, 5)
	fmt.Println(got)
	if got {
		t.Fail()
	}

	got = checkDiagonalCross(path, 1, 1, 2, 2)
	fmt.Println(got)
	if got {
		t.Fail()
	}
}

func TestMain(m *testing.M) {

	os.Exit(m.Run())
}
