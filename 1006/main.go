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

func solve(n, w int, enemies []int) {
	m := make([][16]int, n)
	numPairs := greedyPairs(0, n, w, enemies, 0, m)
	answer := numPairs + (n-numPairs)*2
	fmt.Println(answer)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func greedyPairs(i, n, w int, enemies []int, shadow int, m [][16]int) int {
	// fast returns
	if i == n {
		return 0
	}
	if m[i][shadow] != 0 {
		return m[i][shadow] - 1
	}

	var (
		r  = (i + 1) % n // right
		b  = i + n       // bottom
		br = r + n       // bottom right
	)

	nextShadow := func(leftBits, rightBits int) int {
		if i == 0 {
			leftBits <<= 2
		} else {
			leftBits = shadow & 12
		}
		return leftBits | rightBits
	}

	// assume no link
	x := greedyPairs(i+1, n, w, enemies, nextShadow(0, 0), m)

	// assume left link
	if shadow&3 == 0 && enemies[i]+enemies[b] <= w {
		here := 1
		next := greedyPairs(i+1, n, w, enemies, nextShadow(3, 0), m)
		x = max(x, here+next)
	}

	both := 0

	// assume top link
	if shadow&1 == 0 && i != r && enemies[i]+enemies[r] <= w {
		if !(i == n-1 && shadow&4 != 0) {
			here := 1
			next := greedyPairs(i+1, n, w, enemies, nextShadow(1, 1), m)
			x = max(x, here+next)
			both++
		}
	}

	// assume bottom link
	if shadow&2 == 0 && b != br && enemies[b]+enemies[br] <= w {
		if !(i == n-1 && shadow&8 != 0) {
			here := 1
			next := greedyPairs(i+1, n, w, enemies, nextShadow(2, 2), m)
			x = max(x, here+next)
			both++
		}
	}

	// assume top & bottom link
	if both == 2 {
		here := 2
		next := greedyPairs(i+1, n, w, enemies, nextShadow(3, 3), m)
		x = max(x, here+next)
	}

	// cache it
	m[i][shadow] = x + 1
	return x
}
