/*
 * @Author: hongwei.sun
 * @Date: 2024-04-03 19:21:36
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-03 19:21:37
 * @Description: file content
 */
// https://leetcode.com/problems/trapping-rain-water/
/*
Given n non-negative integers representing an elevation map where the width of each bar is 1,
compute how much water it is able to trap after raining.


The above elevation map is represented by array [0,1,0,2,1,0,1,3,2,1,2,1].
In this case, 6 units of rain water (blue section) are being trapped. Thanks Marcos for contributing this image!

Example:

Input: [0,1,0,2,1,0,1,3,2,1,2,1]
Output: 6
*/
// 这是一个非常 common 的问题 amazon facebook google 都是高频题

/*
整体的思路是这样的：
一个位置能存的水量 取决于它左边的最大值和右边的最大值还有当前位置的值 这个想一下就能理解
所以解决这个问题其实就是找每个位置i  min(maxleft，maxright) - height[i] then sum them
所以directly的想法是维持三个数组， maxleft  maxright and current
tc O(n), SC O(n)

use two pointers to reduce complexity of SC.

TC: O(n)
SC: O(1)
*/
func trap(height []int) int {
	res := 0
	left := 0
	right := len(height) - 1
	leftMax, rightMax := height[left], height[right]
	for left < right {
		if leftMax < rightMax {
			// becasue letfMax is smaller than rightMax already
			// although we don't know the real max of right in this position, 
			// but we know min leftMax and rightMax must be letfMax
			// so we use leftMax minus height[left]
			h := leftMax - height[left]
			if h > 0 {
				res += h
			}
			left++
			// update leftMax for next position
            leftMax = max(leftMax, height[left])
		} else {
			// becasue rightMax is smaller than leftMax already
			// although we don't know the real max of left in this position, 
			// but we know min leftMax and rightMax must be rightMax
			// so we use rightMax minus height[right]
			h := rightMax - height[right]
			if h > 0 {
				res += h
			}
			right--
            rightMax = max(rightMax, height[right])
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

/*
https://leetcode.com/problems/trapping-rain-water/solution/
维护一个非递增栈, 当有一个高是大于queue中的 top 时, 说明这个queue 的 top 的位置可以存水了
可存的水是pop 出这个 top 之后 current 的 index 和剩下的 top 之间可以存的水
所以高度是 min * distance
TC: O(n)
SC: O(n)
*/
func trap(height []int) int {
	res := 0
	index := 0
	queue := make([]int, 0)
	for index < len(height) {
		for len(queue) != 0 && height[index] > height[queue[len(queue)-1]] {
			top := queue[len(queue)-1]
			queue = queue[:(len(queue) - 1)]
			// when the second index, the first index couldn't store water
			if len(queue) == 0 {
				break
			}
			distance := index - queue[len(queue)-1] - 1
			boundedHeight := min(height[index], height[queue[len(queue)-1]]) - height[top]
			res += boundedHeight * distance
		}
		queue = append(queue, index)
		index++
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}