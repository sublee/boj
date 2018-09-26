package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)

func nextInt() int {
	var n int
	_, err := fmt.Fscan(reader, &n)

	if err != nil {
		panic(err)
	}

	return n
}

func main() {
	t := nextInt()

	for i := 0; i < t; i++ {
		x := nextInt()
		y := nextInt()
		fmt.Println(solve(x, y))
	}
}

func solve(x, y int) int {
	return _solve(y-x, 0)
}

const wrong = -1

func min(a, b, c int) int {
	min := a
	if min == wrong || b != wrong && b < min {
		min = b
	}
	if min == wrong || c != wrong && c < min {
		min = c
	}
	return min
}

func _solve(d, k int) int {
	if d <= 0 {
		// too far
		return wrong
	}

	if d == 1 {
		// before the goal
		if k <= 2 {
			return 1
		}

		// too fast
		return wrong
	}

	walk := func(k int) int {
		if k <= 0 {
			return wrong
		}
		return _solve(d-k, k)
	}
	a := walk(k)
	b := walk(k + 1)
	c := walk(k - 1)

	return 1 + min(a, b, c)
}
