package main

import (
	"testing"
)

func TestDraw(t *testing.T) {
	drawToImage(100, 100, 10, []Path{{points: []Point[int]{{1, 1}, {2, 2}}}}, "test")
}
