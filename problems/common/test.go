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
	res := productExceptSelf([]int{1, 2, 3, 4})
	fmt.Println(res)
}

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	prefix := make([]int, len(nums))
	suffix := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		product := 1
		if i != 0 {
			product = prefix[i-1] * nums[i-1]
		}
		prefix[i] = product
	}
	fmt.Println(prefix)
	for i := len(nums) - 1; i >= 0; i-- {
		product := 1
		if i != len(nums)-1 {
			product = suffix[i+1] * nums[i+1]
		}
		suffix[i] = product
	}
	fmt.Println(suffix)
	for i := 0; i < len(nums); i++ {
		res[i] = prefix[i] * suffix[i]
	}

	return res
}
