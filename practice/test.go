/*
 * @Author: hongwei.sun
 * @Date: 2024-04-03 14:54:08
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-03 22:23:20
 * @Description: file content
 */
func findMin(nums []int) int {
	l := 0
	r := len(nums)-1
	res := nums[0]
	for l <= r {
		if nums[l] < nums[r] {
			res = min(res, nums[l])
			break
		}
		mid := l + (r-l)/2
		res = min(res, nums[mid])
		if nums[mid] > nums[l] {
			l = mid+1
		} else {
			r = mid-1
		}
	}
	return res
}

 func max(a, b int) int {
	if a > b {
		return a
	}
	return b
 }

 func min(a, b int) int {
	if a < b {
		return a
	}
	return b
 }
