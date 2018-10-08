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
	stop := gcm(m, n)

	for {
		if x < y {
			x += m
		} else if x > y {
			y += n
		}

		if x == y {
			return x
		}
		if x > stop {
			return -1
		}
	}
}

func gcm(m, n int) int {
	bigger := m
	if n > m {
		bigger = n
	}

	for {
		if bigger%m == 0 && bigger%n == 0 {
			return bigger
		}
		bigger++
	}
}
