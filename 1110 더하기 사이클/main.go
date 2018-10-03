package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	var i int
	x := n
	for i = 1; ; i++ {
		x = transform(x)
		if x == n {
			break
		}
	}

	fmt.Println(i)
}

func transform(n int) int {
	var a, b int

	// a
	a = n % 10

	// b
	x := 0
	for n != 0 {
		x += n % 10
		n /= 10
	}
	b = x % 10

	return a*10 + b
}
