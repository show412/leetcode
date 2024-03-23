/*
 * @Author: hongwei.sun
 * @Date: 2021-01-22 18:45:51
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-03-23 23:13:40
 * @Description: file content
 */
// https://leetcode.com/problems/largest-rectangle-in-histogram/
/*
Given n non-negative integers representing the histogram's bar height where the width of each bar is 1, find the area of largest rectangle in the histogram.




Above is a histogram where width of each bar is 1, given height = [2,1,5,6,2,3].




The largest rectangle is shown in the shaded area, which has area = 10 unit.



Example:

Input: [2,1,5,6,2,3]
Output: 10
*/
/*
维护一个非递减的栈
当一个元素进来的时候, 如果它小于当前栈的最后一个元素，意味着最后一个元素不能继续向右扩展了
这时候pop出栈中最后一个元素, 计算max area。for stack 直到要进栈的元素大于栈中最后一个元素
为了计算宽度，这时候这个进栈元素的index应该是pop出的最后一个元素的index。
当遍历完整个数组后, 再遍历整个栈，要计算的宽度应该是整个heights数组的长度减去当前的index：
因为这时候已经没有元素了，在栈中的元素的宽度要算上既能向左(也就是它当时能pop出代替的元素)
也能向右也就是heights长度
*/
func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	area := 0
	// 0 store index, 1 store height
	stack := make([][2]int, 0)
	
	for i := 0; i <= len(heights)-1; i++ {
		index := i
		for len(stack) != 0 && heights[i] < stack[len(stack)-1][1] {
			cur := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			/* 
			notice, here is i - cur[0], 
			because when it pop, the last index in stack cur[0] is index
			it's one difficult point
			 */
			area = max(area, (i-cur[0])*cur[1])
			index = cur[0]
		}
		stack = append(stack, [2]int{index, heights[i]})
	}
	// when end of traversing heights, if stack is not blank
	// we need to continue to caucualte left items area in stack
	// notice here is to use heights length minus index in stack
	for _, item := range stack {
		area = max(area, (len(heights)-item[0])*item[1])
	}
	return area
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
