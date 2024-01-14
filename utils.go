package main

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
