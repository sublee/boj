/*
https://www.acmicpc.net/problem/1000
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 두 정수 A와 B를 입력받은 다음,
// A+B를 출력하는 프로그램을 작성하시오.
func main() {
	// 입력:
	// 첫째 줄에 A와 B가 주어진다. (0 < A, B < 10)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	parts := strings.SplitN(input, " ", 2)

	// 출력:
	// 첫째 줄에 A+B를 출력한다.
	a, _ := strconv.Atoi(parts[0])
	b, _ := strconv.Atoi(parts[1])

	fmt.Println(a + b)
}
