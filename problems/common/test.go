package main

import (
	"fmt"
	"strconv"
	// "math"
)

/*
test case:
Input: nums = [0, 1, 3, 50, 75], lower = 0 and upper = 99,
Output: ["2", "4->49", "51->74", "76->99"]


[]
1
1

output: ["1"]
*/
func main() {
	res := findMissingRanges([]int{}, 1, 3)
	fmt.Println(res)
}

func findMissingRanges(nums []int, lower int, upper int) []string {
	res := make([]string, 0)
	if len(nums) == 0 {
		res = append(res, formatRange(lower, upper))
		return res
	}
	if lower < nums[0] {
		res = append(res, formatRange(lower, nums[0]-1))
	}
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1]-nums[i] > 1 {
			res = append(res, formatRange(nums[i]+1, nums[i+1]-1))
		}
	}
	if nums[len(nums)-1] < upper {
		res = append(res, formatRange(nums[len(nums)-1]+1, upper))
	}
	return res
}

func formatRange(s int, e int) string {
	str := strconv.Itoa(s) + "->" + strconv.Itoa(e)
	if s == e {
		str = strconv.Itoa(s)
	}
	return str
}
