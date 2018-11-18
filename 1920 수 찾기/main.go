package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscan(r, &n)

	a := make(map[int]bool)
	for i := 0; i < n; i++ {
		var x int
		fmt.Fscan(r, &x)
		a[x] = true
	}

	var m int
	fmt.Fscan(r, &m)

	for i := 0; i < m; i++ {
		var x int
		fmt.Fscan(r, &x)

		if a[x] {
			fmt.Fprintln(w, 1)
		} else {
			fmt.Fprintln(w, 0)
		}
	}
	w.Flush()
}
