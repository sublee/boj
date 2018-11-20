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

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &a[i])
	}

	fmt.Println(solve(a))
}

func solve(a []int) int {
	m := make(map[key]int)
	return f(a, m, 0, -1, false)
}

type key struct {
	r, i int
}

func f(a []int, m map[key]int, r, i int, choose bool) int {
	if x, ok := m[key{r, i}]; ok {
		return x
	}

	n := len(a)

	var chosen, skipped int

	if i < n-1 {
		if r < 2 {
			chosen = f(a, m, r+1, i+1, true)
		}
		skipped = f(a, m, 0, i+1, false)
	}

	x := chosen
	if skipped > x {
		x = skipped
	}

	if choose {
		x += a[i]
	}

	m[key{r, i}] = x
	return x
}
