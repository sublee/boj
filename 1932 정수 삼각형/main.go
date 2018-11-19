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
	pyramid := make([][]int, n)

	for i := 0; i < n; i++ {
		a := make([]int, i+1)
		pyramid[i] = a
		for j := 0; j < i+1; j++ {
			fmt.Fscan(r, &a[j])
		}
	}

	fmt.Println(solve(pyramid))
}

func solve(pyramid [][]int) int {
	m := make([][]int, len(pyramid))
	return _solve(pyramid, 0, 0, m)
}

func _solve(pyramid [][]int, i, j int, m [][]int) int {
	if len(pyramid) == i {
		return 0
	}

	if len(m[i]) == 0 {
		m[i] = make([]int, i+1)
	} else if m[i][j] != 0 {
		return m[i][j] - 1
	}

	a := _solve(pyramid, i+1, j, m)
	b := _solve(pyramid, i+1, j+1, m)
	if a < b {
		a = b
	}

	x := pyramid[i][j] + a
	// 0 in m means not cached.
	m[i][j] = x + 1
	return x
}
