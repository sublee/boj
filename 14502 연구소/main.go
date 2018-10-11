package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(r, &n, &m)

	lab := make([][]int, n)
	for i := 0; i < n; i++ {
		lab[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(r, &lab[i][j])
		}
	}

	fmt.Println(solve(lab, n, m))
}

func score(lab [][]int, n, m int) int {
	lab2 := _copyLab(lab, n, m)
	_infect(lab2, n, m, 0, 0)

	score := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if lab2[i][j] == 0 {
				score++
			}
		}
	}
	return score
}

func _copyLab(lab [][]int, n, m int) [][]int {
	lab2 := make([][]int, n)
	for i := 0; i < n; i++ {
		lab2[i] = make([]int, m)
		copy(lab2[i], lab[i])
	}
	return lab2
}

func _infect(lab [][]int, n, m int, i, j int) {
	if lab[i][j] == 2 {
		// north
		if i > 0 && lab[i-1][j] == 0 {
			lab[i-1][j] = 2
			_infect(lab, n, m, i-1, j)
		}
		// south
		if i < n-1 && lab[i+1][j] == 0 {
			lab[i+1][j] = 2
			_infect(lab, n, m, i+1, j)
		}
		// west
		if j > 0 && lab[i][j-1] == 0 {
			lab[i][j-1] = 2
			_infect(lab, n, m, i, j-1)
		}
		// east
		if j < m-1 && lab[i][j+1] == 0 {
			lab[i][j+1] = 2
			_infect(lab, n, m, i, j+1)
		}
	}

	j++
	if j == m {
		i++
		j = 0
	}
	if i == n {
		return
	}

	_infect(lab, n, m, i, j)
}

func solve(lab [][]int, n, m int) int {
	return _solve(lab, n, m, 0, 0, 3)
}

func _solve(lab [][]int, n, m int, i, j, c int) int {
	if c == 0 {
		return score(lab, n, m)
	}
	if i == n {
		// Solution not found.
		return -1
	}

	nextI, nextJ := i, j+1
	if nextJ == m {
		nextI, nextJ = i+1, 0
	}

	max := _solve(lab, n, m, nextI, nextJ, c)

	if lab[i][j] == 0 {
		lab[i][j] = 1

		s := _solve(lab, n, m, nextI, nextJ, c-1)
		if s > max {
			max = s
		}

		lab[i][j] = 0
	}

	return max
}
