package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(solve(n))
}

func solve(n int) int {
	if n < 10 {
		return n
	}
	if n > 1022 {
		// There's no answer beyond 9876543210.
		return -1
	}

	var m, m2 [10][]int

	i := 0
	for ; i < 10; i++ {
		m[i] = []int{i}
	}

	place := 10

	for {
		for j := 1; j < 10; j++ {
			nums := []int{}

			for k := 0; k < j; k++ {
				for _, n := range m[k] {
					nums = append(nums, j*place+n)
				}
			}

			if n < i+len(nums) {
				return nums[n-i]
			}

			m2[j] = nums
			i += len(nums)
		}

		place *= 10
		m = m2
		m2 = [10][]int{}
	}
}
