package main

import (
	"fmt"
)

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

	grid := initGrid(gridWidth, gridHeight)

	paths := computePathsOnGrid(gridWidth/2, gridHeight/2, 3, grid, output)

	drawToImage(gridWidth, gridHeight, size, paths, "out")
}
