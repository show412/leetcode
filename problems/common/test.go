package main

import (
	"fmt"
)

/*
test case:
[0,0]
-1
true

[0,1,0]
0
false

[23, 2, 6, 4, 7],  k=0
false

[5,0,0]
0
true

{23, 2, 6, 4, 7}, 6
true

[23, 2, 4, 6, 7], 6
true
*/
func main() {
	res := maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	fmt.Println(res)
}

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	maxSum := nums[0]
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if max <= 0 {
			max = num
		} else {
			max += num
		}
		if max > maxSum {
			maxSum = max
		}
	}
	return maxSum
}
