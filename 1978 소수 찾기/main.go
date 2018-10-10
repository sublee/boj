package main

import "fmt"

func main() {
	coefs := make([]int, 1000)

	var n int
	fmt.Scan(&n)

	var x, counter int
	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		if aks(x, coefs) {
			counter++
		}
	}

	fmt.Println(counter)
}

// https://www.geeksforgeeks.org/aks-primality-test/
func aks(n int, coefs []int) bool {
	if n < 2 {
		return false
	}

	coef(n, coefs)

	coefs[0]++
	coefs[n]--

	i := n
	for i > -1 && coefs[i]%n == 0 {
		i--
	}

	return i < 0
}

func coef(n int, coefs []int) {
	coefs[0] = 1

	for i := 0; i < n; i++ {
		coefs[i+1] = 1

		for j := i; j > 0; j-- {
			coefs[j] = coefs[j-1] - coefs[j]
		}

		coefs[0] = -coefs[0]
	}
}
