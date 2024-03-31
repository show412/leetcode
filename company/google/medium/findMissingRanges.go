import (
	"fmt"
	"math"
	"strconv"
)

// https://leetcode.com/problems/missing-ranges/
/*
Given a sorted integer array nums,
where the range of elements are in the inclusive range [lower, upper],
return its missing ranges.

Example:

Input: nums = [0, 1, 3, 50, 75], lower = 0 and upper = 99,
Output: ["2", "4->49", "51->74", "76->99"]
*/
/*
test case:
Input: nums = [0, 1, 3, 50, 75], lower = 0 and upper = 99,
Output: ["2", "4->49", "51->74", "76->99"]


[]
1
1

output: ["1"]
*/
// 从nums 出发 然后考虑lower和upper在nums里的数的关系 我怎么就想不到呢
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

// 这个算法是说去遍历lower到upper 然后通过nums把他们分割 不是一个好的算法 并且没有考虑一些edge case
func findMissingRanges(nums []int, lower int, upper int) []string {
	point := 0
	limit := math.MaxInt32
	lastLimit := math.MaxInt32
	res := make([]string, 0)
	for i := lower; i <= upper+1; i++ {
		if (point < len(nums) && i == nums[point]) || i == upper+1 {
			lastLimit = limit
			limit = i
			fmt.Println(lastLimit)
			fmt.Println(limit)
			if limit != math.MaxInt32 && lastLimit != math.MaxInt32 && (limit-lastLimit-1) >= 1 {
				rangeFrom := lastLimit + 1
				rangeTo := limit - 1
				str := strconv.Itoa(rangeFrom) + "->" + strconv.Itoa(rangeTo)
				if rangeFrom == rangeTo {
					str = strconv.Itoa(rangeFrom)
				}
				res = append(res, str)
			}
			point++
		}
	}
	return res
}
