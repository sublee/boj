package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var c rune
	a := make([]int, 10)

	for i := 0; i < 10; i++ {
		n, _ := fmt.Fscanf(r, "%c", &c)
		if n == 0 || c < '0' || c > '9' {
			a = a[:i]
			break
		}
		a[i] = int(c - '0')
	}

	sort.Sort(sort.Reverse(sort.IntSlice(a)))

	for _, x := range a {
		fmt.Fprint(w, x)
	}
	fmt.Fprintln(w)
	w.Flush()
}
