package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	var t int
	fmt.Fscanln(r, &t)

	var a, b int
	var buf bytes.Buffer

	for i := 0; i < t; i++ {
		fmt.Fscanln(r, &a, &b)
		fmt.Fprintln(&buf, a+b)
	}

	fmt.Print(buf.String())
}
