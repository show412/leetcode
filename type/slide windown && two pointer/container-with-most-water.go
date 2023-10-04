// https://leetcode.com/problems/container-with-most-water/
/*
Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai). n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together with x-axis forms a container, such that the container contains the most water.

Note: You may not slant the container and n is at least 2.





The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.



Example:

Input: [1,8,6,2,5,4,8,3,7]
Output: 49
*/
// 用双指针 不能用单调栈 因为这是算最大面积
// store water depends on the smallest side
// move the pointer with smaller height
func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	area := 0
	for l != r {
		if height[l] < height[r] {
			area = max(area, (r-l)*height[l])
			l++
		} else {
			area = max(area, (r-l)*height[r])
			r--
		}
	}
	return area
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
