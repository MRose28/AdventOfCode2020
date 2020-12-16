package main

import (
	"fmt"
	"mrose.de/aoc/2020/day16"
	"time"
)

//Solve the puzzle
func main() {
	defer elapsed("calculation")()
	fmt.Printf("%d\n\n",day16.Solve())
}

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}
