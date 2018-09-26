package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, x int
	var a, b []int

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		a = append(a, x)
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		b = append(b, x)
	}

	fmt.Println(solve(n, a, b))
}

func solve(n int, a, b []int) int {
	sort.Ints(a)
	sort.Sort(sort.Reverse(sort.IntSlice(b)))

	s := 0
	for i := range a {
		s += a[i] * b[i]
	}

	return s
}
