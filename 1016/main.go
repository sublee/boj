package main

import (
	"fmt"
	"math"
)

func main() {
	var min, max int64
	fmt.Scan(&min, &max)
	fmt.Println(solve(min, max))
}

func solve(min, max int64) int {
	squares := make(map[int64]bool)

	for i := 2; i <= int(math.Sqrt(float64(max))); i++ {
		iSq := int64(i) * int64(i)

		a := int64(math.Ceil(float64(min) / float64(iSq)))
		b := int64(math.Floor(float64(max) / float64(iSq)))

		for j := a; j <= b; j++ {
			x := iSq * j
			squares[x] = true
		}
	}

	length := int(max - min)
	return length - len(squares) + 1
}
