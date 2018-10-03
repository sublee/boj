package main

import (
	"fmt"
	"math"
)

var table = map[string]int{
	"black":  0,
	"brown":  1,
	"red":    2,
	"orange": 3,
	"yellow": 4,
	"green":  5,
	"blue":   6,
	"violet": 7,
	"grey":   8,
	"white":  9,
}

func pow10(n int) int {
	return int(math.Pow(10, float64(n)))
}

func main() {
	var color1, color2, color3 string

	fmt.Scanln(&color1)
	fmt.Scanln(&color2)
	fmt.Scanln(&color3)

	var (
		tens  = table[color1]
		units = table[color2]
		mul   = pow10(table[color3])
	)
	resistance := (tens*10 + units) * mul

	fmt.Println(resistance)
}
