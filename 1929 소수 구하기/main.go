package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var m, n int
	fmt.Scan(&m, &n)

	// Sieve of Eratosthenes
	composites := make([]bool, n+1)
	composites[0] = true
	composites[1] = true

	for i := 2; i*i <= n; i++ {
		if !composites[i] {
			for j := i * 2; j <= n; j += i {
				composites[j] = true
			}
		}
	}

	// Print prime numbers between m and n.
	w := bufio.NewWriter(os.Stdout)
	for i := m; i <= n; i++ {
		if !composites[i] {
			fmt.Fprintln(w, i)
		}
	}
	w.Flush()
}
