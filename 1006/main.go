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

var counter int

type memo map[int]map[int]int

func hash(pairs [][2]int, covered []int) int {
	if len(covered) == 0 {
		return 0
	}

	x := 0

	for i, pair := range pairs {
		j := i * 2
		if has(covered, pair[0]) {
			x |= 1 << uint(j)
		}
		if has(covered, pair[1]) {
			x |= 1 << uint(j+1)
		}
	}

	return x
}

func greedyPairs(pairs [][2]int, covered []int, m memo) (n int) {
	h := hash(pairs, covered)
	n, cached := m[len(pairs)][h]
	if cached {
		return
	}

	defer func() {
		if _, ok := m[len(pairs)]; !ok {
			m[len(pairs)] = make(map[int]int)
		}
		m[len(pairs)][h] = n
		// fmt.Fprintln(os.Stderr, "f(", len(pairs), ", ", h, ") =>", foundPairs)
	}()

	counter++

	if len(pairs) == 0 {
		n = 0
		return
	}

	pair := pairs[0]
	x := pair[0]
	y := pair[1]

	var neg, pos int

	if has(covered, x) || has(covered, y) {
		n = 0
	} else {
		n = 1
	}

	if len(pairs) == 1 {
		return
	}

	// if this pair is not used.
	neg = greedyPairs(pairs[1:], covered, m)

	// if this pair is used.
	if n == 0 {
		pos = 0
	} else {
		subCovered := append(covered, x, y)
		pos = 1 + greedyPairs(pairs[1:], subCovered, m)
	}

	if neg > pos {
		n = neg
	} else {
		n = pos
	}
	return
}

func solve(n, w int, enemies []int) {
	// fmt.Println("---")
	// covered := make([]bool, len(enemies))
	// count := 0

	// linkMap := make(map[int][]int)
	// var byNumLinks [3][]int

	pairs := pairwise(n, w, enemies)
	// fmt.Println("PAIRS:", pairs)

	m := make(memo)
	numPairs := greedyPairs(pairs, make([]int, 0), m)
	answer := numPairs + (n-numPairs)*2

	fmt.Println(answer, counter)
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
