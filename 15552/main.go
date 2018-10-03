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
	fmt.Fscanln(r, &t)

	var a, b int
	for i := 0; i < t; i++ {
		fmt.Fscanln(r, &a, &b)
		fmt.Fprintln(w, a+b)
	}

	w.Flush()
}
