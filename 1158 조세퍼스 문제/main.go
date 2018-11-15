package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	print(solve(n, m))
}

func print(seq []int) {
	w := bufio.NewWriter(os.Stdout)

	fmt.Fprint(w, "<")

	for i, n := range seq {
		if i > 0 {
			fmt.Fprint(w, ", ")
		}
		fmt.Fprint(w, n)
	}

	fmt.Fprintln(w, ">")
	w.Flush()
}

func solve(n, m int) []int {
	ring := make([]int, n)

	for i := 0; i < n; i++ {
		ring[i] = i + 1
	}

	removed := make([]int, n)
	for i := 0; i < n; i++ {
		j := (m - 1) % len(ring)
		removed[i] = ring[j]
		ring = append(ring[j+1:], ring[:j]...)
	}

	return removed
}
