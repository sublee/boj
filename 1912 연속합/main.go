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
	n := len(a)

	sum := a[0]
	max := a[0]
	l, r := 0, 0

	commit := func(l, r int) {
		m := max
		for j := l; j < r; j++ {
			m -= a[j]
			if m > max {
				max = m
			}
		}
	}

	for i := 1; i < n; i++ {
		if sum < 0 {
			commit(l, r)
			l = i
			sum = 0
		}

		sum += a[i]

		if sum > max {
			max = sum
			r = i
		}
	}
	commit(l, r)

	return max
}
