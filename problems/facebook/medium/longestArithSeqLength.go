// https://leetcode.com/problems/longest-arithmetic-sequence/
/*
Given an array A of integers,
return the length of the longest arithmetic subsequence in A.

Recall that a subsequence of
A is a list A[i_1], A[i_2], ..., A[i_k] with 0 <= i_1 < i_2 < ... < i_k <= A.length - 1,
and that a sequence B is arithmetic if B[i+1] - B[i] are all the same value (for 0 <= i < B.length - 1).

Example 1:

Input: [3,6,9,12]
Output: 4
Explanation:
The whole array is an arithmetic sequence with steps of length = 3.
Example 2:

Input: [9,4,7,2,10]
Output: 3
Explanation:
The longest arithmetic subsequence is [4,7,10].
Example 3:

Input: [20,1,15,3,10,5,8]
Output: 4
Explanation:
The longest arithmetic subsequence is [20,15,10,5].

Note:

2 <= A.length <= 2000
0 <= A[i] <= 10000
*/
/*
必须是连续的 不能是跳跃的
比如 17, 43, 52, 78 ： 52-17 = 35 78-43 = 35 但是实际这些数不符合要求
*/
func longestArithSeqLength(A []int) int {
	if len(A) == 2 {
		return 2
	}
	max := 2
	// 在第几个数差值有多少个 这个数据结构是这道题的关键
	dp := make([]map[int]int, len(A))
	// 这里也要从0开始 因为要初始化第0个index的map
	for i := 0; i < len(A); i++ {
		x := A[i]
		// 直接定义map[int]int不能直接赋值 所以需要在这里make才能赋值
		dp[i] = make(map[int]int, 0)
		for j := 0; j < i; j++ {
			y := A[j]
			d := x - y
			if _, ok := dp[j][d]; !ok {
				dp[j][d] = 1
			}
			dp[i][d] = dp[j][d] + 1
			if dp[i][d] > max {
				max = dp[i][d]
			}
		}
	}
	return max
}
