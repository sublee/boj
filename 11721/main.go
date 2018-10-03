package main

import (
	"fmt"
	"os"
)

func main() {
	b := make([]byte, 10)

	for {
		n, _ := os.Stdin.Read(b)
		if n == 0 {
			break
		}
		fmt.Println(string(b[:n]))
	}
}
