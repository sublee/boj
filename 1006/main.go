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

var miss, total int

func f(i, n, w int, enemies []int, shadow int, m [][16]int) int {
	total++

	if i == n {
		// fmt.Println(i, counter, 0)
		return 0
	}

	key := shadow
	if i != n-1 {
		key &^= 12
	}

	if m[i][key] != 0 {
		return m[i][key] - 1
	}

	miss++

	var (
		r  = (i + 1) % n // right
		b  = i + n       // bottom
		br = (i+1)%n + n // bottom right
	)

	nextShadow := func(nextBits, firstBits int) int {
		if i == 0 {
			firstBits <<= 2
		} else {
			firstBits = shadow & 12
		}
		return firstBits | nextBits
	}

	x := f(i+1, n, w, enemies, nextShadow(0, 0), m)

	if shadow&3 == 0 && enemies[i]+enemies[b] <= w {
		x = max(x, 1+f(i+1, n, w, enemies, nextShadow(0, 3), m))
	}

	both := 0

	if shadow&1 == 0 && enemies[i]+enemies[r] <= w {
		if !(i == n-1 && shadow&4 != 0) {
			x = max(x, 1+f(i+1, n, w, enemies, nextShadow(1, 1), m))
			both++
		}
	}

	if shadow&2 == 0 && enemies[b]+enemies[br] <= w {
		if !(i == n-1 && shadow&8 != 0) {
			x = max(x, 1+f(i+1, n, w, enemies, nextShadow(2, 2), m))
			both++
		}
	}

	if both == 2 {
		x = max(x, 2+f(i+1, n, w, enemies, nextShadow(3, 3), m))
	}

	m[i][key] = x + 1
	return x
}

func solve(n, w int, enemies []int) {
	m := make([][16]int, n)
	numPairs := f(0, n, w, enemies, 0, m)
	answer := numPairs + (n-numPairs)*2
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
