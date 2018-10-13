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

	stack := []int{}
	cmds := []bool{}
	i := 0

	for len(a) != 0 {
		top := len(stack) - 1

		if top < 0 || stack[top] < a[0] {
			i++
			stack = append(stack, i)
			cmds = append(cmds, true)
		} else if stack[top] == a[0] {
			a = a[1:]
			stack = stack[:top]
			cmds = append(cmds, false)
		} else {
			break
		}
	}

	if len(a) != 0 {
		fmt.Fprintln(w, "NO")
	} else {
		for _, x := range cmds {
			if x {
				fmt.Fprintln(w, "+")
			} else {
				fmt.Fprintln(w, "-")
			}
		}
	}

	w.Flush()
}
