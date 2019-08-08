// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
/*
Given an array of integers nums sorted in ascending order, find the starting and ending position of a given target value.

Your algorithm's runtime complexity must be in the order of O(log n).

If the target is not found in the array, return [-1, -1].

Example 1:

Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]
Example 2:

Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]
*/
// it is a common solution to binary search to find range of target
// two binary search to find the boundary left and right
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	res := make([]int, 2)
	start := 0
	end := len(nums) - 1
	// find the boundary left
	for start+1 < end {
		mid := start + (end-start)/2
		if nums[mid] > target {
			end = mid
		} else if nums[mid] == target {
			// because we need to find the left boundary, so it needs to find the small range
			// so end = mid there
			/*
				这时候可能有这样一个疑问 为什么要end = mid 如果mid前还有target的怎么办
				因为要找范围小的， 所以就应该向前缩小范围，即使找到了mid == target的 也没有return
				还有下一次循环， 直到start+1 >= end 这个时候 end就是最小的target所在的位置
				找右边的边界同理 就是向后缩小范围
				有没有可能错过一个target 不在start和end的范围内？
				不可能，因为一直是等于的时候 end往前提 不可能遗漏一个target
			*/
			end = mid
		} else {
			start = mid
		}
	}
	// 因为要找最左范围 注意这里先是start
	if nums[start] == target {
		res[0] = start
	} else if nums[end] == target {
		res[0] = end
	} else {
		return []int{-1, -1}
	}

	// find the boundary right
	start = 0
	end = len(nums) - 1
	for start+1 < end {
		mid := start + (end-start)/2
		if nums[mid] > target {
			end = mid
		} else if nums[mid] == target {
			// because we need to find the right boundary,
			// so it needs to find the big range
			// so start = mid there
			start = mid
		} else {
			start = mid
		}
	}
	// 因为要找最右范围 注意这里先是end
	if nums[end] == target {
		res[1] = end
	} else if nums[start] == target {
		res[1] = start
	} else {
		return []int{-1, -1}
	}
	return res
}
