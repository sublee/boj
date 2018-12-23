package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	var a, b []int

	fmt.Scan(&n)

	a = make([]int, n)
	b = make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	copy(b, a)
	sort.Ints(b)

	order := make(map[int][]int, n)
	for i, x := range b {
		order[x] = append(order[x], i)
	}

	for _, x := range a {
		o := order[x][0]
		order[x] = order[x][1:]
		fmt.Print(o)
		fmt.Print(" ")
	}
}
