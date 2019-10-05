package main

import (
	"fmt"
)

/*
test case:
[]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3
10

[]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2
6
*/
func main() {
	res := maxProduct([]int{2, 3, -2, 4})
	// 998001
	fmt.Println(res)
}

func maxProduct(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	res, imax, imin := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if num < 0 {
			// when num is smaller than 0, max multiply it will be more smaller
			// so swap max and min
			imax, imin = imin, imax
		}
		imax = max(num, num*imax)
		imin = min(num, num*imin)
		res = max(res, imax)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
