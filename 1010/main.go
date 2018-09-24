package main

import "fmt"

func main() {
	var t int
	var n, m int

	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		fmt.Scan(&n, &m)
		fmt.Println(solve(n, m))
	}
}

func solve(n, m int) int {
	if n == 1 {
		return m
	}

	r := 1

	// Fix bridges one by one.
	for i := m - 1; i >= n; i-- {
		r += memoize(solve)(n-1, i)
	}

	return r
}

// Memoization

type pair struct {
	n, m int
}

var memo = make(map[pair]int)

func memoize(f func(int, int) int) func(int, int) int {
	return func(n, m int) int {
		if r, ok := memo[pair{n, m}]; ok {
			return r
		}

		r := f(n, m)
		memo[pair{n, m}] = r
		return r
	}
}
