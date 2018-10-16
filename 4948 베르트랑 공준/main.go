package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var n, max int
	var ns []int
	for {
		fmt.Fscan(r, &n)
		if n == 0 {
			break
		}
		ns = append(ns, n)
		if max < n {
			max = n
		}
	}

	primes := eratosthenes(max * 2)

	for _, n := range ns {
		count := 0
		for i := n + 1; i <= n*2; i++ {
			if primes[i] {
				count++
			}
		}
		fmt.Fprintln(w, count)
	}
	w.Flush()
}

func eratosthenes(n int) []bool {
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

	primes := composites
	for i, x := range primes {
		primes[i] = !x
	}

	return primes
}
