import "math"

// https://leetcode.com/problems/split-array-largest-sum/
/*
Given an array which consists of non-negative integers and an integer m,
you can split the array into m non-empty continuous subarrays.
Write an algorithm to minimize the largest sum among these m subarrays.

Note:
If n is the length of array, assume the following constraints are satisfied:

1 ≤ n ≤ 1000
1 ≤ m ≤ min(50, n)
Examples:

Input:
nums = [7,2,5,10,8]
m = 2

Output:
18

Explanation:
There are four ways to split nums into two subarrays.
The best way is to split it into [7,2,5] and [10,8],
where the largest sum among the two subarrays is only 18.
*/
// DP to solve this problem
func splitArray(nums []int, m int) int {
	// sub[i] means the first i items sum
	sub := make([]int, len(nums)+1)
	sub[0] = 0
	for i := 0; i < len(nums); i++ {
		sub[i+1] = sub[i] + nums[i]
	}
	// f[i][j] means the max sum that i items is divided into j parts
	f := make([][]int, len(nums)+1)
	for i := 0; i < len(nums)+1; i++ {
		f[i] = make([]int, m+1)
	}
	for i := 0; i < len(nums)+1; i++ {
		for j := 0; j < m+1; j++ {
			f[i][j] = math.MaxInt32
		}
	}
	f[0][0] = 0
	for i := 1; i <= len(nums); i++ {
		for j := 1; j <= m; j++ {
			for k := 0; k < i; k++ {
				/*
					this is the key of this problem
					find out the max between first k combine j-1 parts and the left nums(i-k)
					then findout the min between these results
				*/
				f[i][j] = min(f[i][j], max(f[k][j-1], sub[i]-sub[k]))
			}
		}
	}
	return f[len(nums)][m]
}
func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
Intuition

The problem satisfies the non-aftereffect property. We can try to use dynamic programming to solve it.

The non-aftereffect property means, once the state of a certain stage is determined, it is not affected by the state in the future. In this problem, if we get the largest subarray sum for splitting nums[0..i] into j parts, this value will not be affected by how we split the remaining part of nums.

To know more about non-aftereffect property, this link may be helpful : http://www.programering.com/a/MDOzUzMwATM.html

Algorithm

Let's define f[i][j] to be the minimum largest subarray sum for splitting nums[0..i] into j parts.

Consider the jth subarray. We can split the array from a smaller index k to i to form it. Thus f[i][j] can be derived from max(f[k][j - 1], nums[k + 1] + ... + nums[i]). For all valid index k, f[i][j] should choose the minimum value of the above formula.

The final answer should be f[n][m], where n is the size of the array.

For corner situations, all the invalid f[i][j] should be assigned with INFINITY, and f[0][0] should be initialized with 0.


*/
