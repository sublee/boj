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
	sum := a[0]
	max := a[0]
	l, r := 0, 0

	max2 := max

	for i := 1; i < len(a); i++ {
		if sum < 0 {
			m := findMax(a, max, l, r)
			if m > max2 {
				max2 = m
			}
			l = i
			sum = 0
		}

		sum += a[i]

		if sum > max {
			max = sum
			r = i
		}
	}
	m := findMax(a, max, l, r)
	if m > max2 {
		max2 = m
	}

	return max2
}

func findMax(a []int, max, l, r int) int {
	max2 := max

	for i := l; i < r; i++ {
		max -= a[i]
		if max > max2 {
			max2 = max
		}
	}

	return max2
}
