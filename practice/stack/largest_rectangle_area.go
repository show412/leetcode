// https://leetcode.com/problems/largest-rectangle-in-histogram/
/*
Given n non-negative integers representing the histogram's bar height where the width of each bar is 1, find the area of largest rectangle in the histogram.




Above is a histogram where width of each bar is 1, given height = [2,1,5,6,2,3].




The largest rectangle is shown in the shaded area, which has area = 10 unit.



Example:

Input: [2,1,5,6,2,3]
Output: 10
*/
// 这个题有很多细节
func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	area := 0
	// 存的是数组里的index
	stack := make([]int, 0)
	// 这里只所以是<= it's to make the left column in the stack to be caculated
	for i := 0; i <= len(heights); i++ {
		curHeight := -1
		if i != len(heights) {
			curHeight = heights[i]
		}
		for len(stack) != 0 && curHeight < heights[stack[len(stack)-1]] {
			index := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			height := heights[index]
			wide := i
			if len(stack) != 0 {
				wide = i - stack[len(stack)-1] - 1
			}
			area = max(area, wide*height)
		}
		// current column must to be into the stack
		// and the valuse in the stack are not bigger than current column
		stack = append(stack, i)
	}
	return area
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
