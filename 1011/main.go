package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t, x, y int

	r := bufio.NewReader(os.Stdin)

	fmt.Fscan(r, &t)

	for i := 0; i < t; i++ {
		fmt.Fscan(r, &x, &y)
		fmt.Println(solve(x, y))
	}
}

func solve(x, y int) int {
	d := y - x

	// Find maximum 2*k.
	var k2 int
	for d >= k2+2 {
		k2 += 2
		d -= k2
	}

	// 0 or 1 or 2 more steps required depending on the remaining d value.
	k := k2 / 2
	more := (d + k) / (k + 1)

	return k2 + more
}
