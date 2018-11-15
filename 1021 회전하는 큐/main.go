package main

import "fmt"

func main() {
	var n, m int

	fmt.Scan(&n, &m)
	a := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&a[i])
	}

	fmt.Println(solve(n, a))
}

func solve(n int, a []int) int {
	q := make([]int, n)
	for i := 0; i < n; i++ {
		q[i] = i + 1
	}

	rotated := 0

	for len(a) != 0 {
		if a[0] != q[0] {
			n = len(q)

			// Find a[0] in q.
			i := 0
			for ; q[i] != a[0]; i++ {
			}

			// Rotate q.
			q = append(q[i:], q[:i]...)

			if i <= n/2 {
				// Rotated left.
				rotated += i
			} else {
				// Rotated right.
				rotated += n - i
			}
		}

		// Pop.
		a = a[1:]
		q = q[1:]
	}

	return rotated
}
