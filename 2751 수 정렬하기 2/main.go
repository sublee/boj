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

	mergeSort(a, 0, n)

	for i := 0; i < n; i++ {
		fmt.Fprintln(w, a[i])
	}
	w.Flush()
}

func mergeSort(a []int, l, h int) {
	n := h - l
	if n <= 1 {
		return
	}
	m := (l + h) / 2

	mergeSort(a, l, m)
	mergeSort(a, m, h)

	b := make([]int, n)
	i, j := l, m

	for k := 0; k < n; k++ {
		if i == m {
			copy(b[k:], a[j:h])
			break
		} else if j == h {
			copy(b[k:], a[i:m])
			break
		}

		if a[i] < a[j] {
			b[k] = a[i]
			i++
		} else {
			b[k] = a[j]
			j++
		}
	}

	copy(a[l:h], b)
}
