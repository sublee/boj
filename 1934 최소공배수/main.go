package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var t int
	fmt.Fscan(r, &t)

	var a, b int
	for i := 0; i < t; i++ {
		fmt.Fscan(r, &a, &b)
		fmt.Fprintln(w, lcm(a, b))
	}
	w.Flush()
}

func lcm(a, b int) int {
	ab := a * b
	if ab < 0 {
		ab = -ab
	}
	return ab / gcd(a, b)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
