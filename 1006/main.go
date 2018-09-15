/*
https://www.acmicpc.net/problem/1006
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var input string

	scanner := bufio.NewScanner(os.Stdin)

	// 첫째 줄에 테스트 케이스의 개수 T가 주어진다.
	scanner.Scan()
	input = scanner.Text()
	t, _ := strconv.Atoi(input)

	for i := 0; i < t; i++ {
		// 첫째 줄에는 (구역의 개수)/2 값 N과
		// 특수 소대원의 수 W가 주어진다.
		// (1 ≤ N ≤ 10000, 1 ≤ W ≤ 10000).
		scanner.Scan()
		input = scanner.Text()
		parts := strings.SplitN(input, " ", 2)
		n, _ := strconv.Atoi(parts[0])
		w, _ := strconv.Atoi(parts[1])

		// 둘째 줄에는 1~N번째 구역에 배치된 적의 수가 주어지고,
		// 셋째 줄에는 N+1 ~ 2N번째 구역에 배치된 적의 수가
		// 공백으로 구분되어 주어진다.
		var enemies []int
		for j := 0; j < 2; j++ {
			scanner.Scan()
			input = scanner.Text()
			parts := strings.SplitN(input, " ", n)
			for k := 0; k < n; k++ {
				x, _ := strconv.Atoi(parts[k])
				enemies = append(enemies, x)
			}
		}

		solve(n, w, enemies)
	}
}

var total int

func f(i, n, w int, enemies, counter []int) int {
	total++

	if i == n*2 {
		return 0
	}

	var (
		// l int // left
		r int // right
		b int // bottom
	)

	if i < n {
		r = (i + 1) % n
		b = i + n
	} else {
		r = (i+1)%n + n
		b = -1
	}

	x := f(i+1, n, w, enemies, counter)
	// fmt.Println(" ", i, "-", x)

	if counter[i] == 0 {
		counter[i]++
		for _, j := range [2]int{r, b} {
			if j == -1 || j == i {
				continue
			}
			if counter[j] != 0 {
				continue
			}

			if enemies[i]+enemies[j] <= w {
				counter[j]++

				y := 1 + f(i+1, n, w, enemies, counter, m)
				// fmt.Println(" ", i, j, y)
				if x < y {
					x = y
				}

				counter[j]--
			}
		}
		counter[i]--
	}

	return x
}

func solve(n, w int, enemies []int) {
	counter := make([]int, n*2)
	numPairs := f(0, n, w, enemies, counter)
	answer := numPairs + (n-numPairs)*2
	fmt.Println(answer, total)
}

// Set operations

func has(a []int, x int) bool {
	i := sort.SearchInts(a, x)
	return i < len(a) && a[i] == x
}

func update(a []int, xs ...int) []int {
	for _, x := range xs {
		i := sort.SearchInts(a, x)
		a = append(a[:i], append([]int{x}, a[i:]...)...)
	}
	return a
}

func hasPair(a [][2]int, p [2]int) bool {
	for _, q := range a {
		if q == p {
			return true
		}
	}
	return false
}

func addPair(a [][2]int, p [2]int) ([][2]int, bool) {
	if hasPair(a, p) {
		return a, false
	}
	return append(a, p), true
}
