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
2， 堆排序  Time complexity: O(nlogk) Space complexity: O(n)
3, 动态规划  Time complexity: O(n) Space complexity: O(n)
refer to https://leetcode.com/articles/sliding-window-maximum/
*/
func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	if n*k == 0 {
		return []int{}
	}
	if k == 1 {
		return nums
	}
	left := make([]int, n)
	left[0] = nums[0]
	right := make([]int, n)
	right[n-1] = nums[n-1]
	for i := 1; i < n; i++ {
		// from left to right
		if i%k == 0 {
			left[i] = nums[i]
		} else {
			left[i] = max(left[i-1], nums[i])
		}
		// from right to left
		j := n - i - 1
		if (j+1)%k == 0 {
			right[j] = nums[j]
		} else {
			right[j] = max(right[j+1], nums[j])
		}
	}
	output := make([]int, n-k+1)
	for i := 0; i < n-k+1; i++ {
		output[i] = max(left[i+k-1], right[i])
	}
	return output
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 这是直接遍历的算法
class Solution {
    public int[] maxSlidingWindow(int[] nums, int k) {
        int n = nums.length;
        if (n * k == 0) return new int[0];

        int [] output = new int[n - k + 1];
        for (int i = 0; i < n - k + 1; i++) {
            int max = Integer.MIN_VALUE;
            for(int j = i; j < i + k; j++)
                max = Math.max(max, nums[j]);
            output[i] = max;
        }
        return output;
    }
}