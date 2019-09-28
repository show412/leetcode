// https://leetcode.com/problems/partition-equal-subset-sum/
/*
Given a non-empty array containing only positive integers, find if the array can be partitioned into two subsets such that the sum of elements in both subsets is equal.

Note:

Each of the array element will not exceed 100.
The array size will not exceed 200.


Example 1:

Input: [1, 5, 11, 5]

Output: true

Explanation: The array can be partitioned as [1, 5, 5] and [11].


Example 2:

Input: [1, 2, 3, 5]

Output: false

Explanation: The array cannot be partitioned into equal sum subsets.
*/
// pre sum is not right, for example [1,2,3,4,5,6,7] could split into [1,2,5,6] and [3,4,7]
func canPartition(nums []int) bool {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	// sum is odd, return false
	if sum%2 == 1 {
		return false
	}
	sum = sum / 2
	// dp[i] define that if or not there is nums could combine into i
	// 其中dp[i]表示原数组是否可以取出若干个数字，其和为i。
	dp := make([]bool, sum+1)
	dp[0] = true
	// state function
	/*
		这种状态转移方程
		如果 dp[j - nums[i]] 为true的话，说明现在已经可以组成 j-nums[i] 这个数字了，
		再加上nums[i]，就可以组成数字j了，那么dp[j]就一定为true。
		如果之前dp[j]已经为true了，当然还要保持true，所以还要‘或’。
		用sum到num的循环的数组的传递来实现在dp中”若干个“的定义实现
	*/
	for _, num := range nums {
		for i := sum; i >= num; i-- {
			dp[i] = dp[i] || dp[i-num]
		}
	}
	return dp[sum]
}
