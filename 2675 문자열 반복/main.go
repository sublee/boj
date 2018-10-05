package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f := bufio.NewReader(os.Stdin)

	var t int
	fmt.Fscanln(f, &t)

	var r int
	var s string

	for i := 0; i < t; i++ {
		fmt.Fscanln(f, &r, &s)

		for _, c := range s {
			fmt.Print(strings.Repeat(string(c), r))
		}
		fmt.Println()
	}
}
