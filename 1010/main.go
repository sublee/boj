package main

import "fmt"

func main() {
	var t int
	var n, m int

	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		fmt.Scan(&n, &m)
		fmt.Println(solve(n, m))
	}
}

func solve(n, m int) int {
	return 1
}
