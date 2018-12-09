package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	var m, n int
	fmt.Fscan(r, &m, &n)

	a := newMatrix(m, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Fscan(r, &a[i][j])
		}
	}

	fmt.Println(solve(m, n, a))
}

func newMatrix(m, n int) [][]int {
	a := make([][]int, m)
	for i := 0; i < m; i++ {
		a[i] = make([]int, n)
	}
	return a
}

type ij struct{ i, j int }

func solve(m, n int, a [][]int) int {
	memo := newMatrix(m, n)
	return _solve(m, n, a, 0, 0, memo)
}

func _solve(m, n int, a [][]int, i, j int, memo [][]int) int {
	if memo[i][j] != 0 {
		// return a cached answer
		return memo[i][j] - 1
	}

	if i == m-1 && j == n-1 {
		// reached the goal
		return 1
	}

	routes := 0
	defer func() {
		memo[i][j] = routes + 1
	}()

	for _, d := range []ij{{-1, 0}, {0, -1}, {0, +1}, {+1, 0}} {
		ni, nj := i+d.i, j+d.j

		if ni == -1 || ni == m || nj == -1 || nj == n {
			// out of range
			continue
		}

		if a[i][j] > a[ni][nj] {
			routes += _solve(m, n, a, ni, nj, memo)
		}
	}

	return routes
}
