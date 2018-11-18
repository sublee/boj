package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(r, &n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &a[i])
	}

	fmt.Println(solve(a))
}

func solve(a []int) int {
	m := make([]int, len(a))
	max := a[0]

	sum := 0
	for i := 0; i < len(a); i++ {
		sum += a[i]
		m[i] = sum

		if sum > max {
			max = sum
		}
	}

	for i := 1; i < len(a); i++ {
		for j := i; j < len(a); j++ {
			sum = m[j] - a[i-1]
			m[j] = sum

			if sum > max {
				max = sum
			}
		}
	}

	return max
}
