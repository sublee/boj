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

	k := k2 / 2

	// Remaining d must be less than 2*k.
	if d == 0 {
		return k2
	} else if d <= k+1 {
		// 1 more step required for 1 ~ k+1.
		return k2 + 1
	} else {
		// 2 more steps required for k+1 ~ 2k.
		return k2 + 2
	}
}
