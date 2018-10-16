package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(fib(n))
}

var m = make(map[int]int)

func fib(n int) int {
	if n <= 1 {
		return n
	}
	if m[n] != 0 {
		return m[n]
	}
	m[n] = fib(n-1) + fib(n-2)
	return m[n]
}
