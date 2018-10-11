package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	fmt.Println(solve(n, k))
}

const inf = (1 << 63) - 1

func solve(n, k int) int {
	i := 0
	min := inf
	try := func(x int) {
		if x < min {
			min = x
		}
	}

	q := []int{n}
	p := []int{} // next queue
	v := make(map[int]bool)

	for len(q) != 0 {
		for len(q) != 0 {
			n, q = q[0], q[1:]

			if n < 0 || n > 100000 {
				continue
			}

			if v[n] {
				continue
			}
			v[n] = true

			if n == k {
				try(i)
				continue
			} else if n > k {
				try(i + (n - k))
				continue
			}

			p = append(p, n*2, n+1, n-1)
		}

		q, p = p, q
		i++
	}

	return min
}
