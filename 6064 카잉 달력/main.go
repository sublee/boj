package main

import "fmt"

func main() {
	var t, m, n, x, y int

	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		fmt.Scan(&m, &n, &x, &y)
		fmt.Println(solve(m, n, x, y))
	}
}

func solve(m, n, x, y int) int {
	stop := m * n // instead of gcd(m, n)

	for {
		if x == y {
			return x
		} else if x > stop {
			return -1
		} else if x < y {
			x += m
		} else if x > y {
			y += n
		}
	}
}
