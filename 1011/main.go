package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)

func nextInt() int {
	var n int
	_, err := fmt.Fscan(reader, &n)

	if err != nil {
		panic(err)
	}

	return n
}

func main() {
	t := nextInt()

	for i := 0; i < t; i++ {
		x := nextInt()
		y := nextInt()
		fmt.Println(solve(x, y))
	}
}

func solve(x, y int) int {
	return 0
}
