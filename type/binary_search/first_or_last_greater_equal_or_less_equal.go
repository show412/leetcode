/*
 * @Author: hongwei.sun
 * @Date: 2024-04-10 13:48:42
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-10 13:53:05
 * @Description: file content
 */
 // 找到第⼀个⼤于等于 target 
 func searchFirstGreaterElement(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high - low)/2
		if nums[mid] >= target {
		// 找到第⼀个⼤于等于 target 
			if (mid == 0) || (nums[mid-1] < target) { 
				return mid
			}
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
	
// ⼆分查找最后⼀个⼩于等于 target 的元素
func searchLastLessElement(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high - low)/2
		if nums[mid] <= target {
			if (mid == len(nums)-1) || (nums[mid+1] > target) { 
				// 找到最后⼀个⼩于等于target 的元素
				return mid
			}
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
