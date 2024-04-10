/*
 * @Author: hongwei.sun
 * @Date: 2024-04-10 13:23:12
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-10 13:39:17
 * @Description: file content
 */
//  https://leetcode.com/problems/find-target-indices-after-sorting-array/
/*
You are given a 0-indexed integer array nums and a target element target.

A target index is an index i such that nums[i] == target.

Return a list of the target indices of nums after sorting nums in non-decreasing order. If there are no target indices, return an empty list. The returned list must be sorted in increasing order.

 

Example 1:

Input: nums = [1,2,5,2,3], target = 2
Output: [1,2]
Explanation: After sorting, nums is [1,2,2,3,5].
The indices where nums[i] == 2 are 1 and 2.
Example 2:

Input: nums = [1,2,5,2,3], target = 3
Output: [3]
Explanation: After sorting, nums is [1,2,2,3,5].
The index where nums[i] == 3 is 3.
Example 3:

Input: nums = [1,2,5,2,3], target = 5
Output: [4]
Explanation: After sorting, nums is [1,2,2,3,5].
The index where nums[i] == 5 is 4.
 

Constraints:

1 <= nums.length <= 100
1 <= nums[i], target <= 100
*/
// find first index equal to target with binary search
func targetIndices(nums []int, target int) []int {
	sort.Ints(nums)
	res := make([]int,0)
	l := 0
	r := len(nums) - 1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	first := -1
	if l< len(nums) && nums[l] == target {
		first = l
	} else if r >= 0  && nums[r] == target {
		first = r
	}
    if first >= 0 {
        for i := first; i < len(nums); i++ {
            if nums[i] != target {
                break
            }
            res = append(res, i)
	    }
    }
	
	return res
}