package main

import (
	"fmt"
)

func main() {
	// [[0, 30],[5, 10],[15, 20]]
	// [[9,10],[4,9],[4,17]]
	res := subarraySum([]int{1, 2, 3}, 3)
	// res := subarraySum([]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 0)
	fmt.Println(res)
}

func subarraySum(nums []int, k int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == k {
			res++
		}
		if i == len(nums)-1 {
			// if nums[i] == k {
			// 	res++
			// }
			break
		}
		target := k - nums[i]
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == target {
				res++
			}
			target -= nums[j]
		}
	}
	return res
}
