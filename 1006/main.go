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

//    ┌──┐  ┌─────┐  ┌─────┐
//  70│60│55│43 57│60│44 50│
// ───┤  │  └──┬──┴──┼──┬──┘
//  58│40│47 90│45 52│80│40
// ───┴──┘     └─────┘  └───

func best(i, n, w int, enemies []int, covered []bool) int {
	var (
		l int // left
		r int // right
		v int // vertical
	)

	l = (i + n - 1) % n
	r = (i + 1) % n

	if i < n {
		v = i + n
	} else {
		l += n
		r += n
		v = i - n
	}

	lx := enemies[i] + enemies[l]
	rx := enemies[i] + enemies[r]
	vx := enemies[i] + enemies[v]

	j := -1
	x := enemies[i]

	if !covered[l] && lx <= w && lx > x {
		j = l
		x = lx
	}

	if !covered[r] && rx <= w && rx > x {
		j = r
		x = rx
	}

	if !covered[v] && vx <= w && vx > x {
		j = v
		x = vx
	}

	return j
}

func solve(n, w int, enemies []int) {
	covered := make([]bool, len(enemies))
	count := 0

	for i := 0; i < n*2; i++ {
		if covered[i] {
			continue
		}

		j := best(i, n, w, enemies, covered)

		if j == -1 {
			continue
		}

		if best(j, n, w, enemies, covered) == i {
			covered[i] = true
			covered[j] = true
			count++
		}
	}

	for _, c := range covered {
		if c == false {
			count++
		}
	}

	fmt.Println(count)
}
