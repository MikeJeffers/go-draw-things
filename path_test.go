package main

import (
	"fmt"
	"testing"
)

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

func TestPathCrossing(t *testing.T) {
	paths := []Path{{points: []Point[int]{{1, 1}, {0, 0}}}}
	if checkAllPathsForCrossing(paths, 0, 0, 1, 0) {
		t.Errorf("Paths should be non crossing")
	}
	if !checkAllPathsForCrossing(paths, 1, 0, 0, 1) {
		t.Errorf("Paths should be crossing")
	}
}

func TestInitGrid(t *testing.T) {
	g := initGrid(10, 10)
	if len(g) != 10 || len(g[0]) != 10 {
		t.Errorf("Unexpected matrix dimensions")
	}
}

func TestGetOffset(t *testing.T) {
	x, y := getOffset(0)
	if x != 0 || y != 1 {
		t.Errorf("Unexpected offset x %d y %d", x, y)
	}
	x, y = getOffset(1)
	if x != 1 || y != 1 {
		t.Errorf("Unexpected offset x %d y %d", x, y)
	}
	x, y = getOffset(7)
	if x != -1 || y != 1 {
		t.Errorf("Unexpected offset x %d y %d", x, y)
	}
	// Test wrapping
	x, y = getOffset(8)
	if x != 0 || y != 1 {
		t.Errorf("Unexpected offset x %d y %d", x, y)
	}
	x, y = getOffset(9)
	if x != 1 || y != 1 {
		t.Errorf("Unexpected offset x %d y %d", x, y)
	}
	// Test negative wrapping
	x, y = getOffset(-1)
	if x != -1 || y != 1 {
		t.Errorf("Unexpected offset x %d y %d", x, y)
	}
}

func TestIsDiagonal(t *testing.T) {
	orthoPairs := []Point[int]{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
		{0, -100},
	}
	diagPairs := []Point[int]{
		{1, -1},
		{-1, 1},
		{1, 1},
		{-1, -1},
		{-100, -100},
	}
	for _, pt := range orthoPairs {
		if isMovementDiagonal(pt.x, pt.y) {
			t.Errorf("Movement should not be diagonal for offsets x %d y %d", pt.x, pt.y)
		}
	}
	for _, pt := range diagPairs {
		if !isMovementDiagonal(pt.x, pt.y) {
			t.Errorf("Movement should be diagonal for offsets x %d y %d", pt.x, pt.y)
		}
	}

}

func TestComputePathBasic(t *testing.T) {
	paths := computePathsOnGrid(10, 10, 0, initGrid(20, 20), "CCC")
	fmt.Println(paths)
	if len(paths) < 1 {
		t.Errorf("Unexpected path output")
	} else if len(paths[0].points) < 3 {
		t.Errorf("Unexpected path output")
	}
}

func TestComputePath(t *testing.T) {
	paths := computePathsOnGrid(10, 10, 1, initGrid(20, 20), "AABCCABAAABBCCBBCCCBBA")
	fmt.Println(paths)
	if len(paths) < 1 {
		t.Errorf("Unexpected path output")
	} else if len(paths[0].points) < 3 {
		t.Errorf("Unexpected path output")
	}
}
