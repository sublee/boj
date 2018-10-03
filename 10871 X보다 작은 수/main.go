package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	var n, x int
	fmt.Fscan(r, &n, &x)

	var y int
	first := true

	for i := 0; i < n; i++ {
		fmt.Fscan(r, &y)

		if y < x {
			if first {
				first = false
			} else {
				fmt.Print(" ")
			}
			fmt.Print(y)
		}
	}
	fmt.Println()
}
