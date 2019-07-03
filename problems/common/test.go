package main

import (
	"fmt"
	"math"
	// "math"
)

func main() {

	res := splitArray([]int{7, 2, 5, 10, 8}, 2)
	fmt.Println(res)
}

func splitArray(nums []int, m int) int {
	// sub[i] means the first i items sum
	sub := make([]int, len(nums)+1)
	sub[0] = 0
	for i := 0; i < len(nums); i++ {
		sub[i+1] = sub[i] + nums[i]
	}
	// f[i][j] means the max sum that i items is divided into j parts
	f := make([][]int, len(nums)+1)
	for i := 0; i < len(nums)+1; i++ {
		f[i] = make([]int, m+1)
	}
	for i := 0; i < len(nums)+1; i++ {
		for j := 0; j < m+1; j++ {
			f[i][j] = math.MaxInt32
		}
	}
	f[0][0] = 0
	for i := 1; i <= len(nums); i++ {
		for j := 1; j <= m; j++ {
			for k := 0; k < i; k++ {
				/*
					this is the key of this problem
					find out the max between first k combine j-1 parts and the left nums(i-k)
					then findout the min between these results
				*/
				f[i][j] = min(f[i][j], max(f[k][j-1], sub[i]-sub[k]))
			}
		}
	}
	return f[len(nums)][m]
}
func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
