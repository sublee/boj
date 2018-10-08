package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscan(r, &n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &a[i])
	}

	countingSort10000(a)

	for i := 0; i < n; i++ {
		fmt.Fprintln(w, a[i])
	}
	w.Flush()
}

func countingSort10000(a []int) {
	n := len(a)
	b := make([]int, n)

	var counter [10000]int

	for _, x := range a {
		counter[x-1]++
	}

	for k := 1; k < 10000; k++ {
		counter[k] += counter[k-1]
	}

	for i := n - 1; i >= 0; i-- {
		x := a[i]
		j := counter[x-1] - 1
		b[j] = x
		counter[x-1]--
	}

	copy(a, b)
}
