/*
 * @Author: hongwei.sun
 * @Date: 2024-03-28 12:36:17
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-03-28 12:47:44
 * @Description: file content
 */
// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/
/*
Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e.,  [0,1,2,4,5,6,7] might become  [4,5,6,7,0,1,2]).

Find the minimum element.

You may assume no duplicate exists in the array.

Example 1:

Input: [3,4,5,1,2]
Output: 1
Example 2:

Input: [4,5,6,7,0,1,2]
Output: 0
*/
/*
some points:
1, 不管怎么rotate， 这个数组肯定分为一部分排好序的 一部分是因为rotate没有排好的
2， 因为被rotate了最小值肯定在没有排好序的数组里
3，  how to find substring unsorted, we can compare mid and nums[l]
if mid > nums[l] means left half is sorted otherwise right left is sorted
4, move l or r unsorted half substring, the unsorted half substring can be done in the same manner
5, notice 
  1)if l is smaller than right, means it's sorted already
  2)we need to compare mid and current res if mid is exact the min value

最小值肯定在unsorted subarray里， 总是去unsorted里去找 然后通过min比较 最后l==r 一定能找到
*/
func findMin(nums []int) int {
	l := 0
	r := len(nums) - 1
	var res = nums[0]
	for l <= r {
		// means current substring is sorted
		if nums[l] < nums[r] {
			// maybe mid is smallest value
			res = min(nums[l], res)
			break
		}
		mid := l + (r-l)/2
		// maybe mid is smallest value
		res = min(nums[mid], res)
		// there needs to be +1 or -1, otherwise it possiblely will dead loop
		if nums[mid] >= nums[l] {
			l = mid + 1 
		} else {
			r = mid - 1
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
