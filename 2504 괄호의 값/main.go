package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(solve(s))
}

type bracket int

const (
	illegal bracket = iota
	round
	square
)

func (b bracket) score() int {
	switch b {
	case round:
		return 2
	case square:
		return 3
	default:
		return 0
	}
}

var brackets = map[byte]bracket{
	'(': round,
	')': round,
	'[': square,
	']': square,
}

func solve(s string) int {
	stack := []bracket{}
	scores := [][]int{}

	for i := 0; i < len(s); i++ {
		depth := len(stack)
		for len(scores) < depth+1 {
			scores = append(scores, []int{})
		}

		b := brackets[s[i]]

		switch s[i] {
		case '(', '[':
			push(&stack, b)
			scores[depth] = append(scores[depth], b.score())

		case ')', ']':
			if top(stack) != b {
				// Unmatched ending bracket.
				return 0
			}
			pop(&stack)
		}

		// Calculate score when popped.
		if len(stack) < depth {
			summary := sum(scores[depth])
			if summary == 0 {
				summary = 1
			}
			scores[depth] = scores[depth][:0]
			_scores := scores[depth-1]
			_scores[len(_scores)-1] *= summary
		}
	}

	if len(stack) != 0 {
		// An open bracket remains.
		return 0
	}

	if len(scores) == 0 {
		// No bracket found.
		return 0
	}

	return sum(scores[0])
}

func push(stack *[]bracket, b bracket) {
	*stack = append(*stack, b)
}

func pop(stack *[]bracket) bracket {
	b := top(*stack)
	if b == illegal {
		return b
	}

	depth := len(*stack)
	*stack = (*stack)[:depth-1]
	return b
}

func top(stack []bracket) bracket {
	depth := len(stack)
	if depth == 0 {
		return illegal
	}
	return stack[depth-1]
}

func sum(ints []int) int {
	total := 0
	for _, n := range ints {
		total += n
	}
	return total
}
