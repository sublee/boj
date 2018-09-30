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

	cost := make([][3]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &cost[0], &cost[1], &cost[2])
	}

	fmt.Println(solve(cost))
}

type memoKey struct {
	prev, i int
}

func solve(cost [][3]int) int {
	m := make(map[memoKey]int)
	return f(cost, 0, -1, m)
}

func f(cost [][3]int, i, prev int, m map[memoKey]int) int {
	if i == len(cost) {
		return 0
	}

	min := -1

	for color := 0; color < 3; color++ {
		if color == prev {
			continue
		}

		x := cost[i][color] + f(cost, i+1, color, m)
		if min == -1 || x < min {
			min = x
		}
	}

	return min
}
