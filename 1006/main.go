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

func pairwise(n, w int, enemies []int) (pairs [][2]int, counter []int) {
	counter = make([]int, n*2)

	var added bool
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

				pairs, added = addPair(pairs, pair)

				if added {
					counter[i]++
					counter[j]++
				}
			}
		}
	}

	return
}

var miss, total int

type tree struct {
	values map[int]int
	index  map[int]tree
}

func newTree() tree {
	return tree{make(map[int]int), make(map[int]tree)}
}

func greedyPairs(pairs [][2]int, counter []int, covered []int, m tree) (n int) {
	total++

	p := len(pairs)

	if p == 0 {
		n = 0
		return
	}

	t := m
	var path []int
	for _, c := range covered {
		if counter[c] == 0 {
			continue
		}
		path = append(path, c)
		if _, ok := t.index[c]; !ok {
			t.index[c] = newTree()
		}
		t = t.index[c]
	}

	// fmt.Println("000___", fmt.Sprintf("%02d", total))
	// fmt.Println("000___", fmt.Sprintf("%02d", total), "counter:", counter)
	// fmt.Println("000___", fmt.Sprintf("%02d", total), "path:", path, "covered:", covered)

	if cached, ok := t.values[p]; ok {
		n = cached
		// fmt.Println("000___", fmt.Sprintf("%02d", total), "HIT!", p, pairs, n)
		return
	}

	defer func(x int) {
		t.values[p] = n
		// fmt.Println("000___", fmt.Sprintf("%02d", x), "MISS", p, pairs, n)
	}(total)

	miss++

	// ------------------

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

	counter[x]--
	counter[y]--

	// if this pair is not used.
	neg = greedyPairs(pairs[1:], counter, covered, m)

	// if this pair is used.
	if n == 0 {
		pos = 0
	} else {
		pos = 1 + greedyPairs(pairs[1:], counter, update(covered, x, y), m)
	}

	counter[x]++
	counter[y]++

	if neg > pos {
		n = neg
	} else {
		n = pos
	}
	return
}

func solve(n, w int, enemies []int) {
	pairs, counter := pairwise(n, w, enemies)

	covered := make([]int, 0)
	m := newTree()

	numPairs := greedyPairs(pairs, counter, covered, m)
	answer := numPairs + (n-numPairs)*2

	// fmt.Println("PAIRS", pairs)
	// fmt.Println("COUNTER", counter)
	fmt.Println(answer, miss, total)
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
