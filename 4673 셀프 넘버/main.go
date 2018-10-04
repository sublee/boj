package main

import "fmt"

func main() {
	m := make(map[int]bool)

	for i := 1; i <= 10000; i++ {
		m[d(i)] = true

		if m[i] {
			delete(m, i)
			continue
		}

		fmt.Println(i)
	}
}

func d(n int) int {
	x := n

	for ; n != 0; n /= 10 {
		x += n % 10
	}

	return x
}
