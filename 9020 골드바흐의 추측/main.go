package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	var t, max int
	fmt.Fscan(r, &t)

	ns := make([]int, t)

	for i := 0; i < t; i++ {
		fmt.Fscan(r, &ns[i])
		if max < ns[i] {
			max = ns[i]
		}
	}

	primes := eratosthenes(max)

	for _, n := range ns {
		a := n / 2
		b := a

		for {
			if !primes[a] {
				a--
				continue
			}
			if !primes[b] {
				b++
				continue
			}

			if a+b == n {
				break
			}

			if n-a < b {
				a--
			} else {
				b++
			}
		}

		fmt.Println(a, b)
	}
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
