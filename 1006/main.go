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

//    ┌──┐  ┌─────┐  ┌─────┐
//  70│60│55│43 57│60│44 50│
// ───┤  │  └──┬──┴──┼──┬──┘
//  58│40│47 90│45 52│80│40
// ───┴──┘     └─────┘  └───

func tryMerge(i, n, w int, enemies []int, covered []bool) []int {
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

	x := enemies[i]
	lx := x + enemies[l]
	rx := x + enemies[r]
	vx := x + enemies[v]

	tab := make(map[int][]int)

	if l != i && !covered[l] && lx <= w {
		tab[lx] = append(tab[lx], l)
	}
	if r != i && !covered[r] && rx <= w {
		tab[rx] = append(tab[rx], r)
	}
	if v != i && !covered[v] && vx <= w {
		tab[vx] = append(tab[vx], v)
	}

	var keys []int
	for key := range tab {
		keys = append(keys, key)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	var links []int
	for _, key := range keys {
		links = append(links, tab[key]...)
	}

	return links
}

func firstNotCovered(a []int, covered []bool) int {
	for _, j := range a {
		if !covered[j] {
			return j
		}
	}
	return -1
}

func countNotCovered(a []int, covered []bool) int {
	count := 0
	for _, j := range a {
		if !covered[j] {
			count++
		}
	}
	return count
}

func solve(n, w int, enemies []int) {
	// fmt.Println("---")
	covered := make([]bool, len(enemies))
	count := 0

	linkMap := make(map[int][]int)
	var byNumLinks [3][]int

	for i := 0; i < n*2; i++ {
		linkMap[i] = tryMerge(i, n, w, enemies, covered)

		numLinks := len(linkMap[i])
		if numLinks != 0 {
			k := numLinks - 1
			byNumLinks[k] = append(byNumLinks[k], i)
		}
	}

	// fmt.Println(linkMap)

	for _, is := range byNumLinks {
		for _, i := range is {
			if covered[i] {
				continue
			}

			links := linkMap[i]

			// find minimum links
			minLinks := 4
			lonelyJ := -1
			for _, j := range links {
				if covered[j] {
					continue
				}

				linksJ := linkMap[j]
				numLinksJ := countNotCovered(linksJ, covered)

				if numLinksJ < minLinks {
					lonelyJ = j
					minLinks = numLinksJ
				}
			}

			if lonelyJ != -1 {
				// fmt.Println(i, lonelyJ)
				covered[i] = true
				covered[lonelyJ] = true
				count++
			}
		}
	}

	for _, c := range covered {
		if c == false {
			count++
		}
	}

	fmt.Println(count)
}
