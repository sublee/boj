package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	counter := solve(a, b, c)

	for _, c := range counter {
		fmt.Println(c)
	}
}

func solve(a, b, c int) [10]int {
	x := a * b * c

	var counter [10]int

	for x != 0 {
		counter[x%10]++
		x /= 10
	}

	return counter
}
