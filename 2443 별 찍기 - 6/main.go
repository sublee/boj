package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Print(strings.Repeat(" ", i))
		fmt.Print(strings.Repeat("*", (n-i)*2-1))
		fmt.Println()
	}
}
