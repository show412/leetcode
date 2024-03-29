// https://leetcode.com/problems/coin-change/
/*
You are given coins of different denominations
and a total amount of money amount.
Write a function to compute the fewest number of coins
that you need to make up that amount.
If that amount of money cannot be made up by any combination of the coins,
return -1.

Example 1:

Input: coins = [1, 2, 5], amount = 11
Output: 3
Explanation: 11 = 5 + 5 + 1

Example 2:

Input: coins = [2], amount = 3
Output: -1
Note:
You may assume that you have an infinite number of each kind of coin.
*/
// it's actually a backpack problem, but it's a infinite backpack problem
// 因为这里coins可以认为是无限制取的 所以初始化状态方程是1维就可以
// refer to explain to https://www.cnblogs.com/grandyang/p/5138186.html
func coinChange(coins []int, amount int) int {
	// f[i] means the minize coins number for i amount
	// it's the key for this problem
	f := make([]int, amount+1)
	for i := 0; i < amount+1; i++ {
		f[i] = amount + 1
	}
	f[0] = 0
	for i := 1; i < amount+1; i++ {
		for j := 0; j < len(coins); j++ {
			if i >= coins[j] {
				f[i] = min(f[i], f[i-coins[j]]+1)
			}
		}
	}
	if f[amount] > amount {
		return -1
	}
	return f[amount]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
