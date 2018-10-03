package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	scores := make([]int, n)
	max := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&scores[i])

		if max < scores[i] {
			max = scores[i]
		}
	}

	total := 0.0
	for i := 0; i < n; i++ {
		total += float64(scores[i]) / float64(max) * 100
	}

	fmt.Printf("%.2f\n", total/float64(n))
}
