package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanln(r, &n)

	stack := stack([]int{})

	var (
		line []byte
		cmd  string
		arg  int
	)
	for i := 0; i < n; i++ {
		line, _, _ = r.ReadLine()

		parts := strings.Split(string(line), " ")
		cmd = parts[0]

		if len(parts) > 1 {
			arg, _ = strconv.Atoi(parts[1])
		} else {
			arg = 0
		}

		n, print := dispatch(&stack, cmd, arg)

		if print {
			fmt.Fprintln(w, n)
		}
	}
	w.Flush()
}

type stack []int

func (s *stack) push(n int) {
	*s = append(*s, n)
}

func (s *stack) pop() int {
	l := len(*s)
	n := (*s)[l-1]
	*s = (*s)[:l-1]
	return n
}

func (s stack) size() int {
	return len(s)
}

func (s stack) empty() int {
	if len(s) == 0 {
		return 1
	}
	return 0
}

func (s stack) top() int {
	l := len(s)
	n := s[l-1]
	return n
}

func dispatch(s *stack, cmd string, arg int) (int, bool) {
	switch cmd {
	case "push":
		s.push(arg)
		return 0, false

	case "pop":
		if s.empty() == 1 {
			return -1, true
		}
		return s.pop(), true

	case "size":
		return s.size(), true

	case "empty":
		return s.empty(), true

	case "top":
		if s.empty() == 1 {
			return -1, true
		}
		return s.top(), true
	}

	return 0, false
}
