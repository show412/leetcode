package main

import (
	"fmt"
	// "math"
)

/*
test case:
"kkkkzrkatkwpkkkktrq"
"bbbbaaaaababaababab"
*/
func main() {
	res := searchRange([]int{5, 7, 7, 8, 8, 10}, 8)
	fmt.Println(res)
}

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	res := make([]int, 2)
	start := 0
	end := len(nums) - 1
	// find the boundary left
	for start+1 < end {
		mid := start + (end-start)/2
		if nums[mid] > target {
			end = mid
		} else if nums[mid] == target {
			// because we need to find the left boundary, so it needs to find the small range
			// so end = mid there
			end = mid
		} else {
			start = mid
		}
	}
	if nums[start] == target {
		res[0] = start
	} else if nums[end] == target {
		res[0] = end
	} else {
		return []int{-1, -1}
	}

	// find the boundary right
	start = 0
	end = len(nums) - 1
	for start+1 < end {
		mid := start + (end-start)/2
		if nums[mid] > target {
			end = mid
		} else if nums[mid] == target {
			// because we need to find the right boundary, so it needs to find the big range
			// so start = mid there
			start = mid
		} else {
			start = mid
		}
	}
	if nums[end] == target {
		res[1] = end
	} else if nums[start] == target {
		res[1] = start
	} else {
		return []int{-1, -1}
	}
	return res
}
