package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	count := 0
	for i := 1; i <= n; i++ {
		if i < 100 {
			count++
			continue
		}

		d := 10
		a := -1
		b := -1
		var fail bool

		for x := i; x != 0; x /= 10 {
			b = a
			a = x % 10

			if b == -1 {
				continue
			}
			if d == 10 {
				d = a - b
			} else if a-b != d {
				fail = true
				break
			}
		}

		if !fail {
			count++
		}
	}

	fmt.Println(count)
}
