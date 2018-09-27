package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var count int
	var ds []int

	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &count)

	ds = make([]int, count)
	for i := 0; i < count; i++ {
		fmt.Fscan(reader, &ds[i])
	}

	fmt.Println(solve(ds))
}

func solve(ds []int) int {
	if len(ds) == 1 {
		// If there's only 1 divisor, it must be a prime.
		return ds[0] * ds[0]
	}

	sort.Ints(ds)

	var n, p int
	n = 1

	i := 0
	j := len(ds) - 1
	for i <= j {
		if i == j {
			p = ds[i]
		} else {
			p = ds[i] * ds[j]
		}

		if n%p != 0 {
			// Multiply the product of big and small divisor.
			n *= p
		}

		i++
		j--
	}

	return n
}
