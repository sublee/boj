package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	for i := 0; ; i++ {
		b, err := r.ReadByte()
		if err != nil || b < '0' || b > '8' {
			break
		}

		n, _ := strconv.ParseInt(string(b), 8, 8)
		bin := strconv.FormatInt(n, 2)

		if i == 0 {
			fmt.Fprint(w, bin)
		} else {
			fmt.Fprintf(w, "%03s", bin)
		}
	}

	fmt.Fprintln(w)
	w.Flush()
}
