/*
 * @Author: hongwei.sun
 * @Date: 2024-03-11 14:19:39
 * @LastEditors: your name
 * @LastEditTime: 2024-03-11 14:19:40
 * @Description: file content
 */
/*
 * @Author: hongwei.sun
 * @Date: 2021-01-22 18:45:51
 * @LastEditors: your name
 * @LastEditTime: 2024-03-11 14:19:21
 * @Description: file content
 */
import "sort"

// https://leetcode.com/problems/3sum/
/*
Given an array nums of n integers, are there elements a, b, c in nums
such that a + b + c = 0?
Find all unique triplets in the array which gives the sum of zero.

Note:

The solution set must not contain duplicate triplets.

Example:

Given array nums = [-1, 0, 1, 2, -1, -4],

A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]
*/
/*
input: [-2,0,3,-1,4,0,3,4,1,1,1,-3,-5,4,0]
output: [[-5,1,4],[-3,-1,4],[-3,0,3],[-2,-1,3],[-2,1,1],[-1,0,1],[0,0,0]]
*/
/*
1, sort
2, from negative value and skip duplicated value
3, 2sum to find other 2 values
4, find the other solution with other combination
*/
func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		l := i + 1
		r := len(nums) - 1
		// two sum to find other 2 values
		target := 0 - nums[i]
		for l < r {
			sum := nums[l] + nums[r]
			if sum > target {
				r--
			} else if sum < target {
				l++
			} else {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				// if there are other combination, cause it's sorted, move l forward
				// move r is same, just we want to move duplicated value
				l++
				for nums[l] == nums[l-1] && l < r {
					l++
				}
			}
		}
	}
	return res
}
