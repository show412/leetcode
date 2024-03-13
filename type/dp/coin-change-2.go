/*
 * @Author: hongwei.sun
 * @Date: 2024-03-13 13:53:53
 * @LastEditors: your name
 * @LastEditTime: 2024-03-13 13:53:54
 * @Description: file content
 */
/*
 * @Author: hongwei.sun
 * @Date: 2021-01-22 18:45:52
 * @LastEditors: your name
 * @LastEditTime: 2024-03-13 13:52:13
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
