package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"

	"github.com/fogleman/gg"
)

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
	switch dir {
	case 0:
		return 0, 1
	case 1:
		return 1, 1
	case 2:
		return 1, 0
	case 3:
		return 1, -1
	case 4:
		return 0, -1
	case 5:
		return -1, -1
	case 6:
		return -1, 0
	case 7:
		return -1, 1
	default:
		return 0, 1
	}
}

func norm(val int, bounds int) int {
	modVal := val % bounds
	if modVal < 0 {
		modVal = bounds + modVal
	}
	return modVal
}

func drawSpawnOnGrid(startX, startY, dir int, grid *[216][384]int, grammar string, dc *gg.Context) {
	for _, v := range strings.Split(grammar, "") {
		if v == "A" {
			dir += 1
		} else if v == "B" {
			dir -= 1
		}

		for i := 0; i < 7; i++ {
			dir = norm(dir, 8)
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
				mag := 10.0
				dc.SetRGB(1, 1, 1)
				dc.DrawLine(mag*0.5+float64(startX)*mag, mag*0.5+float64(startY)*mag, mag*0.5+float64(nextX)*mag, mag*0.5+float64(nextY)*mag)
				dc.Stroke()
				grid[nextY][nextX] = 1
				startX, startY = nextX, nextY
				if rand.Float32() < 0.1 {
					defer drawSpawnOnGrid(startX, startY, dir, grid, grammar, dc)
				}
				break
			}
		}

	}
}

func main() {
	rules := map[string]string{
		"A": "CBC",
		"B": "ACC",
		"C": "AAB",
	}

	output := runLSystem(rules, "A", 5)

	fmt.Println(output)

	var grid [216][384]int

	dc := gg.NewContext(3840, 2160)
	dc.SetRGB(0, 0, 0)
	dc.DrawRectangle(0, 0, 3840, 2160)
	dc.Fill()

	drawSpawnOnGrid(384/2, 216/2, 3, &grid, output, dc)

	num := int(rand.Float32() * 100)
	dc.SavePNG(fmt.Sprint(num, "_out.jpg"))
}
