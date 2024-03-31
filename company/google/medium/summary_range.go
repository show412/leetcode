import "strconv"

// https://leetcode.com/problems/summary-ranges/
/*
Given a sorted integer array without duplicates,
return the summary of its ranges.

Example 1:

Input:  [0,1,2,4,5,7]
Output: ["0->2","4->5","7"]
Explanation: 0,1,2 form a continuous range; 4,5 form a continuous range.
Example 2:

Input:  [0,2,3,4,6,8,9]
Output: ["0","2->4","6","8->9"]
Explanation: 2,3,4 form a continuous range; 8,9 form a continuous range.
*/
// use three pointer start, end, last
func summaryRanges(nums []int) []string {
	// m := make(map[int]bool, 0)
	if len(nums) == 0 {
		return []string{}
	}
	if len(nums) == 1 {
		return []string{strconv.Itoa(nums[0])}
	}
	var res []string
	start := 1
	end := 0
	last := 0
	for start < len(nums) {
		if nums[start] != nums[last]+1 {
			if last-end >= 1 {
				ranges := strconv.Itoa(nums[end]) + "->" + strconv.Itoa(nums[last])
				res = append(res, ranges)
			} else {
				res = append(res, strconv.Itoa(nums[end]))
			}
			end = start
		}
		last = start
		start++
	}
	if end == len(nums)-1 {
		res = append(res, strconv.Itoa(nums[end]))
	} else {
		ranges := strconv.Itoa(nums[end]) + "->" + strconv.Itoa(nums[last])
		res = append(res, ranges)
	}
	return res
}
