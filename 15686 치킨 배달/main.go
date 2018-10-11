package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(r, &n, &m)

	city := make([][]int, n)
	stores := []pair{}

	for i := 0; i < n; i++ {
		city[i] = make([]int, n)

		for j := 0; j < n; j++ {
			fmt.Fscan(r, &city[i][j])

			// Log a store (2).
			if city[i][j] == 2 {
				stores = append(stores, pair{i, j})
			}
		}
	}

	fmt.Println(solve(city, stores, n, m))
}

type pair struct {
	i, j int
}

// http://cavaliercoder.com/blog/optimized-abs-for-int64-in-go.html
func abs(n int) int {
	y := n >> 63
	return (n ^ y) - y
}

func howFar(a, b pair) int {
	return abs(a.i-b.i) + abs(a.j-b.j)
}

func score(city [][]int, stores []pair, n int) int {
	score := 0

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if city[i][j] != 1 {
				continue
			}

			pos := pair{i, j}

			min := (n * n)
			for _, s := range stores {
				d := howFar(pos, s)
				if d < min {
					min = d
				}
			}
			score += min
		}
	}

	return score
}

func choose(stores []pair, m int) [][]pair {
	l := len(stores)
	if l == m {
		return [][]pair{stores}
	}
	if l == 0 || m == 0 {
		return [][]pair{}
	}

	choices := [][]pair{}

	// Skip [0]
	choices = append(choices, choose(stores[1:], m)...)

	// Keep [0]
	if m == 1 {
		choices = append(choices, []pair{stores[0]})
	} else {
		for _, tail := range choose(stores[1:], m-1) {
			choice := []pair{stores[0]}
			choice = append(choice, tail...)
			choices = append(choices, choice)
		}
	}

	return choices
}

func solve(city [][]int, stores []pair, n, m int) int {
	min := n * n * n
	for _, _stores := range choose(stores, m) {
		x := score(city, _stores, n)
		if x < min {
			min = x
		}
	}
	return min
}
