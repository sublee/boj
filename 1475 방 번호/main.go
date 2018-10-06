package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var counter [9]int
	for {
		x := n % 10

		if x == 9 {
			x = 6
		}

		counter[x]++

		n /= 10
		if n == 0 {
			break
		}
	}

	counter[6] = counter[6]/2 + counter[6]%2

	max := 0
	for _, c := range counter {
		if max < c {
			max = c
		}
	}

	fmt.Println(max)
}
