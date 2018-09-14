/*
https://www.acmicpc.net/problem/1006
*/
package main

import (
	"bufio"
	"fmt"
	"os"
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

func pairwise(n, w int, enemies []int) [][2]int {
	links := make([][]int, n*2)

	for i := 0; i < n*2; i++ {
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

		for _, j := range [2]int{r, b} {
			if j != -1 && j != i && enemies[i]+enemies[j] <= w {
				links[i] = add(links[i], j)
				links[j] = add(links[j], i)
			}
		}
	}

	// Collect unique pairs.
	pairs := make([][2]int, 0)

	for i, js := range links {
		for _, j := range js {
			pair := [2]int{i, j}
			// sort
			if i > j {
				pair[0] = j
				pair[1] = i
			}

			pairs = addPair(pairs, pair)
		}
	}

	return pairs
}

var noPairs [][2]int

func greedyPairs(pairs [][2]int, covered []int) [][2]int {
	// fmt.Println("f(", pairs, ", ", covered, ")")
	if len(pairs) == 0 {
		return noPairs
	}

	pair := pairs[0]
	x := pair[0]
	y := pair[1]

	var foundPairs [][2]int
	if has(covered, x) || has(covered, y) {
		foundPairs = noPairs
	} else {
		foundPairs = [][2]int{pair}
	}

	if len(pairs) == 1 {
		return foundPairs
	}

	// if this pair is not used.
	a := greedyPairs(pairs[1:], covered)

	// if this pair is used.
	b := append(foundPairs, greedyPairs(pairs[1:], add(add(covered, x), y))...)

	// indent := strings.Repeat(" ", 12-len(pairs))
	// fmt.Println(indent, covered)
	if len(a) > len(b) {
		// fmt.Println(indent, pair, " NO")
		// fmt.Println(indent, "=>", a)
		// fmt.Println(indent, "  ", b)
		return a
	}
	// fmt.Println(indent, pair, "YES")
	// fmt.Println(indent, "  ", a)
	// fmt.Println(indent, "=>", b)
	return b
}

func solve(n, w int, enemies []int) {
	// fmt.Println("---")
	// covered := make([]bool, len(enemies))
	// count := 0

	// linkMap := make(map[int][]int)
	// var byNumLinks [3][]int

	pairs := pairwise(n, w, enemies)
	// fmt.Println("PAIRS:", pairs)

	uniqPairs := greedyPairs(pairs, make([]int, 0))
	numPairs := len(uniqPairs)
	answer := numPairs + (n-numPairs)*2

	fmt.Println(answer)
}

// Set operations

func has(a []int, i int) bool {
	for _, j := range a {
		if j == i {
			return true
		}
	}
	return false
}

func add(a []int, i int) []int {
	if has(a, i) {
		return a
	}
	return append(a, i)
}

func hasPair(a [][2]int, p [2]int) bool {
	for _, q := range a {
		if q == p {
			return true
		}
	}
	return false
}

func addPair(a [][2]int, p [2]int) [][2]int {
	if hasPair(a, p) {
		return a
	}
	return append(a, p)
}
