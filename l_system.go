package main

import (
	"fmt"
	"strings"
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
