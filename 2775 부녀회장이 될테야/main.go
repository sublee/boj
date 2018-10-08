package main

import "fmt"

func main() {
	var t, k, n int

	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		fmt.Scan(&k, &n)
		fmt.Println(solve(k, n))
	}
}

func solve(k, n int) int {
	k++

	a := make([][]int, k)
	for i := 0; i < k; i++ {
		a[i] = make([]int, n)
		a[i][0] = 1
	}
	for j := 1; j < n; j++ {
		a[0][j] = j + 1
	}

	for j := 1; j < n; j++ {
		for i := 1; i < k; i++ {
			a[i][j] = a[i][j-1] + a[i-1][j]
		}
	}

	return a[k-1][n-1]
}
