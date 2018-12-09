package main

import "fmt"

func main() {
	var n, f int
	fmt.Scan(&n, &f)

	m := (n / 100 * 100) % f
	fmt.Printf("%02d", (f-m)%f)
	fmt.Println()
}
