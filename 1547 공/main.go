package main

import "fmt"

func main() {
	var m int
	fmt.Scan(&m)

	z := 1

	var x, y int
	for i := 0; i < m; i++ {
		fmt.Scan(&x, &y)
		switch z {
		case x:
			z = y
		case y:
			z = x
		}
	}

	fmt.Println(z)
}
