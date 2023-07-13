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

func drawSpawnOnGrid(startX, startY, dir int, grid [][]int, grammar string) []Path {
	paths := make([]Path, 1)
	path := Path{points: []Point[int]{{x: startX, y: startY}}}
	for _, v := range strings.Split(grammar, "") {
		if v == "A" {
			dir += 1
		} else if v == "B" {
			dir -= 1
		}

		for i := 0; i < len(DIRS); i++ {
			dir = norm(dir, len(DIRS))
			offX, offY := getOffset(dir)
			nextX, nextY := startX+offX, startY+offY
			if len(grid) <= nextY || nextY < 0 {
				dir++
			} else if len(grid[0]) <= nextX || nextX < 0 {
				dir++
			} else if grid[nextY][nextX] == 1 {
				dir++
			} else if math.Abs(math.Abs(float64(offX))-math.Abs(float64(offY))) < 0.0001 &&
				(grid[startY][startX+offX] == 1 || grid[startY+offY][startX] == 1) {
				dir++
			} else {
				grid[nextY][nextX] = 1
				startX, startY = nextX, nextY
				path.points = append(path.points, Point[int]{x: startX, y: startY})
				if rand.Float32() < 0.1 {
					morePaths := drawSpawnOnGrid(startX, startY, dir, grid, grammar)
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
	const gridWidth, gridHeight = 30, 20
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

	paths := drawSpawnOnGrid(gridWidth/2, gridHeight/2, 3, grid, output)

	for _, v := range paths {
		drawPath(v, size, dc)
	}

	num := int(rand.Float32() * 100)
	dc.SavePNG(fmt.Sprint(num, "_test.jpg"))
}
