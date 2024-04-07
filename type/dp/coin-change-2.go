/*
 * @Author: hongwei.sun
 * @Date: 2024-03-13 13:53:53
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-07 23:23:59
 * @Description: file content
 */
// https://leetcode.com/problems/coin-change-2/
/*
You are given coins of different denominations and a total amount of money. Write a function to compute the number of combinations that make up that amount. You may assume that you have infinite number of each kind of coin.



Example 1:

Input: amount = 5, coins = [1, 2, 5]
Output: 4
Explanation: there are four ways to make up the amount:
5=5
5=2+2+1
5=2+1+1+1
5=1+1+1+1+1
Example 2:

Input: amount = 3, coins = [2]
Output: 0
Explanation: the amount of 3 cannot be made up just with coins of 2.
Example 3:

Input: amount = 10, coins = [10]
Output: 1


Note:

You can assume that

0 <= amount <= 5000
1 <= coin <= 5000
the number of coins is less than 500
the answer is guaranteed to fit into signed 32-bit integer
*/
/*
coin-change is about minize we can choose from coins.
here actually it's same we can use same thought solution
just it's not minize
*/
/*
这个二维的dp 相对好理解
dp[i][j] 代表前i个硬币能组成总数为j的个数，要注意最后减的时候 因为是前i个 实际取应该是coins[i-1]
*/
func change(amount int, coins []int) int {
	dp := make([][]int, len(coins)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, amount+1)
	}
	dp[0][0] = 1
	for i := 1; i <= len(coins); i++ {
		dp[i][0] = 1
		for j := 1; j <= amount; j++ {
			dp[i][j] = dp[i-1][j]
			if coins[i-1] <= j {
				dp[i][j] += dp[i][j-coins[i-1]]
			}
		}
	} 
	return dp[len(coins)][amount]
}


func change(amount int, coins []int) int {
	// dp array present coins combination selection number for i
	dp := make([]int, amount+1)
	// why assign 1 here, because we want one zero amount, only 1 selection that's no to select any coins
	dp[0] = 1
	// we need loop coins firstly in this case can avoid to generate same combination
	for i := 0; i < len(coins); i++ {
		// j is from c, because smaller than c is no sense to traverse.
		c := coins[i]
		for j := c; j <= amount; j++ {
			dp[j] = dp[j] + dp[j-c]
		}
	}
	return dp[amount]
}
