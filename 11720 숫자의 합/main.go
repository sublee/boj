package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n int
	fmt.Scanln(&n)

	b := make([]byte, 1)
	sum := 0

	for i := 0; i < n; i++ {
		os.Stdin.Read(b)
		x, _ := strconv.Atoi(string(b))
		sum += x
	}

	fmt.Println(sum)
}
