package main

import (
	"fmt"
)

/*
test case:
101, 56, 69, 48, 30
4

[8,85,24,85,69]
4

[73,106,39,6,26,15,30,100,71,35,46,112,6,60,110]
29

[98,60,24,89,84,51,61,96,108,87,68,29,14,11,13,50,13,104,57,8,57,111,92,87,9,59,65,116,56,39,55,11,21,105,57,36,48,93,20,94,35,68,64,41,37,11,50,47,8,9]
439
*/
func main() {
	res := canPartition([]int{1, 5, 11, 5})
	fmt.Println(res)
}

func canPartition(nums []int) bool {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	// sum is odd, return false
	if sum%2 == 1 {
		return false
	}
	sum = sum / 2
	// dp[i] define that if or not there is nums could combine into i
	dp := make([]bool, sum+1)
	dp[0] = true
	// state function
	for _, num := range nums {
		fmt.Println(dp)
		for i := sum; i >= num; i-- {
			dp[i] = dp[i] || dp[i-num]
		}
	}
	return dp[sum]
}
