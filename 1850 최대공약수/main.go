package main

import (
	"fmt"
	"strings"
)

func main() {
	var a, b uint64
	fmt.Scan(&a, &b)
	fmt.Println(strings.Repeat("1", int(gcd(a, b))))
}

func gcd(a, b uint64) uint64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
