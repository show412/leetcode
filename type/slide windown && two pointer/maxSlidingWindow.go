/*
 * @Author: hongwei.sun
 * @Date: 2024-04-02 23:20:35
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-02 23:20:35
 * @Description: file content
 */
// https://leetcode.com/problems/sliding-window-maximum/
/*
Given an array nums, there is a sliding window of size k which is moving from the very left of the array to the very right. You can only see the k numbers in the window. Each time the sliding window moves right by one position. Return the max sliding window.

Example:

Input: nums = [1,3,-1,-3,5,3,6,7], and k = 3
Output: [3,3,5,5,6,7]
Explanation:

Window position                Max
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7
Note:
You may assume k is always valid, 1 ≤ k ≤ input array's size for non-empty array.

Follow up:
Could you solve it in linear time?
*/
/*
三种方法：
1， 直接遍历  Time complexity : O(nk) Space complexity : O(n−k+1)
2， 单调队列 O(n) SC O(k)
单调队列，存当前滑动窗口里的值，并且第一个永远是当前滑动窗口里最大的值
当窗口滑动的时候 而且当这个元素index已经不在滑动窗口里的时候，把队列里的第一个pop
调整单调队列，第一个永远是当前滑动窗口里最大的值，后面小的也要存，因为它可能是后面的最大值
每一个元素都要入dp， 因为它可能是后面的最大值
refer to https://leetcode.com/articles/sliding-window-maximum/
*/

func maxSlidingWindow(nums []int, k int) []int {
    res := make([]int, 0)
	// 单调队列，存当前滑动窗口里的值，并且第一个永远是当前滑动窗口里最大的值
    dq := make([]int, 0)
    for i, v := range nums {
		// 当窗口滑动的时候 而且当这个元素index已经不在滑动窗口里的时候，把队列里的第一个pop
        if i > k-1 && dq[0] <= i-k {
            dq = dq[1:]
        }
		// 调整单调队列，第一个永远是当前滑动窗口里最大的值，后面小的也要存，因为它可能是后面的最大值
        for len(dq) > 0 && v > nums[dq[len(dq)-1]] {
            dq = dq[:len(dq)-1]
        }
		// 每一个元素都要入dp， 因为它可能是后面的最大值
        dq = append(dq, i)
		// 第一个永远是当前滑动窗口里最大的值
        if i >= k-1 {
            res = append(res, nums[dq[0]])
        }
    }
    return res
}


// 这是直接遍历的算法
/*
first, sliding windows total is n-k+1
but maybe TLE
*/
func maxSlidingWindow(nums []int, k int) []int {
    l := len(nums)
	if l * k == 0 {
		return []int{}
	}
	res := make([]int, l-k+1)
	for i := 0; i< l-k+1; i++ {
		maxNum := math.MinInt64
		for j := i; j < i + k; j++ {
			maxNum = max(maxNum, nums[j])
		}
		res[i] = maxNum
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
