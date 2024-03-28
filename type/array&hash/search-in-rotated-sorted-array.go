/*
 * @Author: hongwei.sun
 * @Date: 2024-03-28 12:52:58
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-03-28 17:08:48
 * @Description: file content
 */
//  https://leetcode.com/problems/search-in-rotated-sorted-array/description/
// There is an integer array nums sorted in ascending order (with distinct values).

// Prior to being passed to your function, nums is possibly rotated at an unknown pivot index k (1 <= k < nums.length) such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed). For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3 and become [4,5,6,7,0,1,2].

// Given the array nums after the possible rotation and an integer target, return the index of target if it is in nums, or -1 if it is not in nums.

// You must write an algorithm with O(log n) runtime complexity.

// Example 1:

// Input: nums = [4,5,6,7,0,1,2], target = 0
// Output: 4
// Example 2:

// Input: nums = [4,5,6,7,0,1,2], target = 3
// Output: -1
// Example 3:

// Input: nums = [1], target = 0
// Output: -1
 

// Constraints:

// 1 <= nums.length <= 5000
// -104 <= nums[i] <= 104
// All values of nums are unique.
// nums is an ascending array that is possibly rotated.
// -104 <= target <= 104
/*
target 如果存在， 那一定在某一个sorted array里 这个sorted array可能很小 可能只有一个元素(nums[mid] == target)
我们要做的就是去确定应该去哪个区间去找target
所以有以下情况:
1, rotated array, r > l 这个数组就是sorted 否则就不是sorted的
2, target可能落在一个sorted array里也可能落在一个unsorted array里
结合以上去拆分情况，缩小范围

notice:
1, 二分查找最后两个指针一定落在一个元素上
2, rotated array的mid 左右肯定有一个是sorted 一个是unsorted的
*/
func search(nums []int, target int) int {
    l, r := 0, len(nums) - 1
	for l <=r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		}
		// left is one sorted substring
		if nums[l] <= nums[mid] {
			// target is in right unsorted substring
			if target > nums[mid] || target < nums[l] {
				l = mid + 1
			} else {
				// target is in left sorted substring
				r = mid -1
			}
		// right is one sorted substring
		} else {
			if target > nums[r] || target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	return -1
}