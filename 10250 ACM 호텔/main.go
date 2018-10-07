package main

import "fmt"

func main() {
	var t, h, w, n int

	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		fmt.Scan(&h, &w, &n)
		fmt.Println(solve(h, w, n))
	}
}

func solve(h, w, n int) int {
	var a, b int

	a = n % h
	b = n / h

	if a == 0 {
		a = h
	} else {
		b++
	}

	return a*100 + b
}
