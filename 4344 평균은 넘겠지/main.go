package main

import (
	"fmt"
)

func main() {
	var c int
	fmt.Scan(&c)

	var n, total, avg, above int
	var scores []int

	for i := 0; i < c; i++ {
		fmt.Scan(&n)

		scores = make([]int, n)
		total = 0
		for j := 0; j < n; j++ {
			fmt.Scan(&scores[j])
			total += scores[j]
		}

		avg = total / n
		above = 0
		for _, score := range scores {
			if score > avg {
				above++
			}
		}

		fmt.Printf("%.3f%%\n", float64(above)/float64(n)*100)
	}
}
