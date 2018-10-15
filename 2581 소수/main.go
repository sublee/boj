package main

import "fmt"

func main() {
	var m, n int
	fmt.Scan(&m, &n)

	// Sieve of Eratosthenes
	composites := make([]bool, n+1)
	composites[1] = true

	for i := 2; i*i <= n; i++ {
		if !composites[i] {
			for j := i * 2; j <= n; j += i {
				composites[j] = true
			}
		}
	}

	// Calculate the prime sum and find the first prime.
	var sum, first int

	for i := m; i <= n; i++ {
		if !composites[i] {
			sum += i

			if first == 0 {
				first = i
			}
		}
	}

	if sum != 0 {
		fmt.Println(sum)
		fmt.Println(first)
	} else {
		fmt.Println(-1)
	}
}
