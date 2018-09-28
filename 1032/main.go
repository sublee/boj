package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(r, &n)

	names := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscanln(r, &names[i])
	}

	fmt.Println(solve(names))
}

func solve(names []string) string {
	var buf bytes.Buffer
	var char byte

	length := len(names[0])

	for i := 0; i < length; i++ {
		char = names[0][i]

		for _, name := range names[1:] {
			if name[i] != char {
				buf.WriteRune('?')
				goto Next
			}
		}

		buf.WriteByte(char)

	Next:
	}

	return buf.String()
}
