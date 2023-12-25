package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"

	"github.com/fogleman/gg"
)

type Point[T comparable] struct {
	x, y T
}

type Path struct {
	points []Point[int]
}

var DIRS []Point[int] = []Point[int]{
	{x: 0, y: 1},
	{x: 1, y: 1},
	{x: 1, y: 0},
	{x: 1, y: -1},
	{x: 0, y: -1},
	{x: -1, y: -1},
	{x: -1, y: 0},
	{x: -1, y: 1},
}

func runLSystem(rules map[string]string, seed string, gens uint) string {
	prevGen := []rune(seed)
	for i := 0; i < int(gens); i++ {

		nextGen := make([]string, 0)
		for _, v := range prevGen {
			val, ok := rules[string(v)]
			if ok {
				nextGen = append(nextGen, val)
			} else {
				nextGen = append(nextGen, string(v))
			}
		}
		prevGen = []rune(strings.Join(nextGen, ""))
		fmt.Println(i, string(prevGen))
	}
	return string(prevGen)
}

func getOffset(dir int) (int, int) {
	return DIRS[dir].x, DIRS[dir].y
}

func norm(val int, bounds int) int {
	modVal := val % bounds
	if modVal < 0 {
		modVal = bounds + modVal
	}
	return modVal
}

func scaleAndOffset(x int, magnitude float64) float64 {
	return magnitude*0.5 + float64(x)*magnitude
}

func drawPath(path Path, step float64, dc *gg.Context) {
	dc.SetRGB(1, 1, 1)

	nilPt := Point[int]{x: -1, y: -1}
	prev := Point[int]{x: -1, y: -1}

	for _, v := range path.points {
		if prev != nilPt {
			dc.DrawLine(scaleAndOffset(prev.x, step), scaleAndOffset(prev.y, step), scaleAndOffset(v.x, step), scaleAndOffset(v.y, step))
		}
		prev = v
	}
	dc.Stroke()
}

// TODO function for more efficient drawing of lines
func arePointsLinear(points []Point[int]) bool {
	if len(points) < 3 {
		return true // 2 points are always linear
	}
	prev := points[0]
	var dx, dy, prevDx, prevDy int
	for i, p := range points {
		if i == 0 {
			continue
		} else if i == 1 {
			dx = p.x - prev.x
			dy = p.y - prev.y
			prevDx, prevDy = dx, dy
		}
		dx = p.x - prev.x
		dy = p.y - prev.y
		if dx != prevDx || dx != prevDy {
			return false // if delta in x or y changes, non linear
		}
		prevDx, prevDy = dx, dy
		prev = p
	}
	return true
}

func rotDirection(dirValue, rotation int) int {
	return dirValue + rotation
}

func getRotationDir(symbol string) int {
	switch symbol {
	case "A":
		return 1
	case "B":
		return -1
	default:
		return 0
	}
}

func checkDiagonalCross(path Path, x1, y1, x2, y2 int) bool {
	var prev Point[int]
	for _, p := range path.points {
		if p.x == x1 && p.y == y2 {
			if prev.x == x2 && prev.y == y1 {
				return true
			}
		} else if p.x == x2 && p.y == y1 {
			if prev.x == x1 && prev.y == y2 {
				return true
			}
		}
		prev = p
	}
	return false
}

func checkAllPathsForCrossing(paths []Path, x1, y1, x2, y2 int) bool {
	for _, p := range paths {
		if checkDiagonalCross(p, x1, y1, x2, y2) {
			return true
		}
	}
	return false
}

func isMovementDiagonal(x, y int) bool {
	return math.Abs(math.Abs(float64(x))-math.Abs(float64(y))) < 0.0001
}

func computePathsOnGrid(startX, startY, dir int, grid [][]int, grammar string) []Path {
	paths := make([]Path, 1)
	path := Path{points: []Point[int]{{x: startX, y: startY}}}
	for _, v := range strings.Split(grammar, "") {
		rotation := getRotationDir(v)
		dir = rotDirection(dir, rotation)

		for i := 0; i < len(DIRS); i++ {
			dir = norm(dir, len(DIRS))
			offX, offY := getOffset(dir)
			nextX, nextY := startX+offX, startY+offY
			if len(grid) <= nextY || nextY < 0 {
				dir = rotDirection(dir, rotation)
			} else if len(grid[0]) <= nextX || nextX < 0 {
				dir = rotDirection(dir, rotation)
			} else if grid[nextY][nextX] == 1 {
				dir = rotDirection(dir, rotation)
			} else if isMovementDiagonal(offX, offY) && (grid[startY][nextX] == 1 && grid[nextY][startX] == 1) {
				if checkAllPathsForCrossing(paths, startX, startY, nextX, nextY) {
					dir = rotDirection(dir, rotation)
				}
			} else {
				grid[nextY][nextX] = 1
				startX, startY = nextX, nextY
				path.points = append(path.points, Point[int]{x: startX, y: startY})
				if rand.Float32() < 0.055 {
					morePaths := computePathsOnGrid(startX, startY, dir, grid, grammar)
					paths = append(paths, morePaths...)
				}
				break
			}
		}

	}
	if len(path.points) > 0 {
		paths = append(paths, path)
	}
	return paths
}

func main() {
	const gridWidth, gridHeight = 384, 216
	const size = 10

	rules := map[string]string{
		"A": "CBC",
		"B": "ACC",
		"C": "AAB",
	}

	output := runLSystem(rules, "A", 5)

	fmt.Println(output)

	var grid = [][]int{}
	grid = make([][]int, gridHeight)
	for i := range grid {
		grid[i] = make([]int, gridWidth)
	}

	dc := gg.NewContext(gridWidth*size, gridHeight*size)
	dc.SetRGB(0, 0, 0)
	dc.DrawRectangle(0, 0, gridWidth*size, gridHeight*size)
	dc.Fill()

	paths := computePathsOnGrid(gridWidth/2, gridHeight/2, 3, grid, output)

	for _, v := range paths {
		drawPath(v, size, dc)
	}

	num := int(rand.Float32() * 100)
	dc.SavePNG(fmt.Sprint(num, "_out.jpg"))
}
