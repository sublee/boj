package main

import (
	"fmt"
	"regexp"
	"strings"
)

var re = regexp.MustCompile("c[=-]|dz=|d-|lj|nj|[sz]=")

func main() {
	var word string
	fmt.Scan(&word)

	word = strings.TrimSpace(word)
	word = re.ReplaceAllString(word, ".")

	fmt.Println(len(word))
}
