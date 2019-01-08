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

	a := make([][]int, n+1)
	for i := range a {
		a[i] = make([]int, m+1)
	}

	var x int
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			fmt.Fscan(r, &x)

			a[i][j] = x + a[i][j-1] + a[i-1][j] - a[i-1][j-1]
		}
	}

	var k int
	fmt.Fscan(r, &k)

	var i1, j1, i2, j2 int // i, j, x, y in the quiz.
	for i := 0; i < k; i++ {
		fmt.Fscan(r, &i1, &j1, &i2, &j2)

		sum := a[i2][j2] + a[i1-1][j1-1] - a[i2][j1-1] - a[i1-1][j2]
		fmt.Fprintln(w, sum)
	}

	w.Flush()
}
