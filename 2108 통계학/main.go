package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(r, &n)

	a := make([]int, n)
	total := 0

	for i := 0; i < n; i++ {
		fmt.Fscan(r, &a[i])
		total += a[i]
	}

	countingSort8001(a)

	// 산술평균 : N개의 수들의 합을 N으로 나눈 값
	fmt.Println(int(math.Round(float64(total) / float64(n))))

	// 중앙값 : N개의 수들을 증가하는 순서로 나열했을 경우 그 중앙에 위치하는 값
	fmt.Println(a[len(a)/2])

	// 최빈값 : N개의 수들 중 가장 많이 나타나는 값
	fmt.Println(findMode(a))

	// 범위 : N개의 수들 중 최댓값과 최솟값의 차이
	fmt.Println(a[len(a)-1] - a[0])
}

func countingSort8001(a []int) {
	n := len(a)
	b := make([]int, n)

	var counter [8001]int

	for _, x := range a {
		counter[x+4000]++
	}

	for k := 1; k < 8001; k++ {
		counter[k] += counter[k-1]
	}

	for i := n - 1; i >= 0; i-- {
		x := a[i]
		j := counter[x+4000] - 1
		b[j] = x
		counter[x+4000]--
	}

	copy(a, b)
}

func findMode(a []int) int {
	prevN := a[0]
	counter := 0

	modes := []int{prevN}
	maxCounter := 0

	for _, n := range a[1:] {
		if prevN != n {
			prevN = n
			counter = 0
		} else {
			counter++
		}

		if counter == maxCounter {
			modes = append(modes, n)
		} else if counter > maxCounter {
			maxCounter = counter
			modes = []int{n}
		}
	}

	if len(modes) > 1 {
		return modes[1]
	}
	return modes[0]
}
