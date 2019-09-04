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
/* 维护一个非递增栈, 当有一个高是大于queue中的 top 时, 说明这个queue 的 top 的位置可以存水了
可存的水是pop 出这个 top 之后 current 的 index 和剩下的 top 之间可以存的水
所以高度是 min * distance
*/
func trap(height []int) int {
	res := 0
	index := 0
	queue := make([]int, 0)
	for index < len(height) {
		for len(queue) != 0 && height[index] > height[queue[len(queue)-1]] {
			top := queue[len(queue)-1]
			queue = queue[:(len(queue)-1)]
			// when the second index, the first index couldn't store water
			if len(queue) == 0 {
				break
			}
			distance := index - queue[len(queue)-1]-1
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

int trap(vector<int>& height)
{
    int left = 0, right = height.size() - 1;
    int ans = 0;
    int left_max = 0, right_max = 0;
    while (left < right) {
        // a taller bar exists on left pointer's right side
        if (height[left] < height[right]) {
            height[left] >= left_max ? (left_max = height[left]) : ans += (left_max - height[left]);
            ++left;
        }
        // a taller bar exists on right pointer's left side
        else {
            height[right] >= right_max ? (right_max = height[right]) : ans += (right_max - height[right]);
            --right;
        }
    }
    return ans;
}