package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	var (
		ch          rune
		size        int
		numWords    int
		isLatin     bool
		prevIsLatin bool
	)

	for {
		ch, size, _ = r.ReadRune()

		if size == 0 {
			break
		}

		prevIsLatin = isLatin
		isLatin = ('a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z')

		if !prevIsLatin && isLatin {
			numWords++
		}
	}

	fmt.Println(numWords)
}
