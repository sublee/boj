package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var i int
	for i = 0; 1+(i*(i+1)/2)*6 < n; i++ {
		//        └─ sum of 1~i: i(i+1)/2
	}

	fmt.Println(i + 1)
}
