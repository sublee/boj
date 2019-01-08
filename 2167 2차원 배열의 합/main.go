package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var n, m int
	fmt.Fscan(r, &n, &m)

	var x int
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(r, &x)

			a[i][j] = x + at(a, i, j-1) + at(a, i-1, j) - at(a, i-1, j-1)
		}
	}

	var k int
	fmt.Fscan(r, &k)

	var i1, j1, i2, j2 int // i, j, x, y in the quiz.
	for i := 0; i < k; i++ {
		fmt.Fscan(r, &i1, &j1, &i2, &j2)

		i1 -= 2
		j1 -= 2
		i2--
		j2--

		sum := at(a, i2, j2) + at(a, i1, j1) - at(a, i2, j1) - at(a, i1, j2)
		fmt.Fprintln(w, sum)
	}

	w.Flush()
}

func at(a [][]int, i, j int) int {
	if 0 <= i && i < len(a) && 0 <= j && j < len(a[i]) {
		return a[i][j]
	}
	return 0
}
