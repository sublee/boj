package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := make([]int, 3)
	fmt.Scan(&nums[0], &nums[1], &nums[2])
	sort.Ints(nums)
	fmt.Println(nums[1])
}
