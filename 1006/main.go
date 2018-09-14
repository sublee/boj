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

func pairwise(n, w int, enemies []int) [][2]int {
	pairs := make([][2]int, 0)

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
				pair := [2]int{i, j}
				// sort
				if i > j {
					pair[0] = j
					pair[1] = i
				}

				pairs = addPair(pairs, pair)
			}
		}
	}

	return pairs
}

var counter int

type memo map[int]*tree

type tree struct {
	value int
	index map[int]*tree
}

func newTree() *tree {
	return &tree{-1, make(map[int]*tree)}
}

func greedyPairs(pairs [][2]int, covered []int, m memo) (n int) {
	p := len(pairs)
	if _, ok := m[p]; !ok {
		m[p] = newTree()
	}

	t, _ := m[p]
	for _, c := range covered {
		if _, ok := t.index[c]; !ok {
			t.index[c] = newTree()
		}
		t = t.index[c]
	}
	if t.value != -1 {
		n = t.value
		// fmt.Println(p)
		return
	}

	defer func(c string) {
		t.value = n
		fmt.Println(p, c, n)
		// fmt.Println(p, covered, n)
		// if _, ok := m[len(pairs)]; !ok {
		// 	m[len(pairs)] = make(map[string]int)
		// }
		// m[len(pairs)][h] = n
		// fmt.Fprintln(os.Stderr, "f(", len(pairs), ", ", h, ") =>", foundPairs)
	}(fmt.Sprintf("%v", covered))

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
		sort.Sort(sort.IntSlice(subCovered))
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
	pairs := pairwise(n, w, enemies)

	covered := make([]int, 0)
	m := make(memo)

	numPairs := greedyPairs(pairs, covered, m)
	answer := numPairs + (n-numPairs)*2

	fmt.Println("PAIRS", pairs)
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

func remove(a []int, i int) []int {
	for k, j := range a {
		if j == i {
			return append(a[:k], a[k+1:]...)
		}
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

func addPair(a [][2]int, p [2]int) [][2]int {
	if hasPair(a, p) {
		return a
	}
	return append(a, p)
}
