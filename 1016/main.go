package main

import (
	"fmt"
	"math"
)

func main() {
	var min, max int
	fmt.Scan(&min, &max)
	fmt.Println(solve(min, max))
}

func solve(min, max int) int {
	squares := make(map[int]bool)

	for i := 2; i <= int(math.Sqrt(float64(max))); i++ {
		iSq := i * i

		a := int(math.Ceil(float64(min) / float64(iSq)))
		b := int(math.Floor(float64(max) / float64(iSq)))

		for j := a; j <= b; j++ {
			x := iSq * j
			squares[x] = true
		}
	}

	length := max - min
	return length - len(squares) + 1
}
