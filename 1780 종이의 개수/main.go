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

	paper := make([][]int, n)
	for i := 0; i < n; i++ {
		paper[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(r, &paper[i][j])
		}
	}

	a := solve(paper, n, 0, 0)

	fmt.Println(a[0])
	fmt.Println(a[1])
	fmt.Println(a[2])
}

func solve(paper [][]int, n, i, j int) [3]int {
	a := [3]int{}

	if n == 1 {
		a[paper[i][j]+1] = 1
		return a
	}

	a = merge(a, solve(paper, n/3, i, j))
	a = merge(a, solve(paper, n/3, i, j+n/3))
	a = merge(a, solve(paper, n/3, i, j+n/3*2))

	a = merge(a, solve(paper, n/3, i+n/3, j))
	a = merge(a, solve(paper, n/3, i+n/3, j+n/3))
	a = merge(a, solve(paper, n/3, i+n/3, j+n/3*2))

	a = merge(a, solve(paper, n/3, i+n/3*2, j))
	a = merge(a, solve(paper, n/3, i+n/3*2, j+n/3))
	a = merge(a, solve(paper, n/3, i+n/3*2, j+n/3*2))

	switch [2]int{0, 0} {

	case [2]int{a[0], a[1]}:
		return [3]int{0, 0, 1}

	case [2]int{a[0], a[2]}:
		return [3]int{0, 1, 0}

	case [2]int{a[1], a[2]}:
		return [3]int{1, 0, 0}

	default:
		return a
	}
}

func merge(a1, a2 [3]int) [3]int {
	a := [3]int{}
	for i := 0; i < 3; i++ {
		a[i] = a1[i] + a2[i]
	}
	return a
}
