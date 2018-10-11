package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	var n, b, c int
	var a []int

	fmt.Fscan(r, &n)

	a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &a[i])
	}

	fmt.Fscan(r, &b, &c)

	fmt.Println(solve(a, b, c))
}

func solve(a []int, b, c int) int {
	counter := len(a)

	for _, x := range a {
		x -= b
		if x <= 0 {
			continue
		}

		counter += x / c
		if x%c != 0 {
			counter++
		}
	}

	return counter
}
