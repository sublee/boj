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
	a = append([]int{0}, a...)
	return f(a, m, 0, 0)
}

type key struct {
	r, i int
}

func f(a []int, m map[key]int, r, i int) int {
	if x, ok := m[key{r, i}]; ok {
		return x
	}

	n := len(a)

	var one, two int

	if r < 2 && i < n-1 {
		one = f(a, m, r+1, i+1)
	}

	if i < n-2 {
		two = f(a, m, 1, i+2)
	}

	if one == 0 && two == 0 && i != n-1 {
		return 0
	}

	max := one
	if two > max {
		max = two
	}

	x := a[i] + max
	m[key{r, i}] = x
	return x
}
