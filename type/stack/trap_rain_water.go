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

/*
we can see height[left] < height[right],
then for pointerleft, he knows a taller bar exists on his right side,
then if leftMax is taller than him,
he can contain some water for sure(in our case).
So we go ans += (left_max - height[left]).
But if leftMax is shorter than him,
then there isn't a left side bar can help him contain water,
then he will become other bars' leftMax.
so execute (left_max = height[left]).
Same idea for right part.
TC: O(n)
SC: O(1)
*/
func trap(height []int) int {
	res := 0
	left := 0
	right := len(height) - 1
	leftMax, rightMax := 0, 0
	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				res += leftMax - height[left]
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				res += rightMax - height[right]
			}
			right--
		}
	}
	return res
}
