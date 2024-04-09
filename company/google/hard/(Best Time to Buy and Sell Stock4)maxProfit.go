/*
 * @Author: hongwei.sun
 * @Date: 2021-01-22 18:45:52
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-09 12:33:44
 * @Description: file content
 */
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv/
/*
Say you have an array for which the ith element is the price of a given stock on day i.

Design an algorithm to find the maximum profit. You may complete at most k transactions.

Note:
You may not engage in multiple transactions at the same time 
(ie, you must sell the stock before you buy again).

Example 1:

Input: [2,4,1], k = 2
Output: 2
Explanation: Buy on day 1 (price = 2) and sell on day 2 (price = 4), profit = 4-2 = 2.
Example 2:

Input: [3,2,6,5,0,3], k = 2
Output: 7
Explanation: Buy on day 2 (price = 2) and sell on day 3 (price = 6), profit = 6-2 = 4.
Then buy on day 5 (price = 0) and sell on day 6 (price = 3), profit = 3-0 = 3.
*/

func maxProfit(k int, prices []int) int {
	if len(prices) <= 1 || k == 0 {
		return 0
	}
	// dp[i][j] means the max profit j days finish i times transaction
	dp := make([][]int, k+1)
	for i := 0; i <= k; i++ {
		dp[i] = make([]int, len(prices))
	}
	/*
		要在在第j天的i次交易的max profit， j天有以下两种情况:
		1, 在j天不交易  max profit = dp[i][j-1]
		2，在j天交易 卖出在某一天买进的， max profit = prices[j] + 当前的最大利润
		如何找在j天的时候的最大利润，看j-1天 可能买了 也可能没有。因为说most k次所以也可能交易小于k
	*/
	for i := 1; i <= k; i++ {
		// 初始化某一天的最大利润，可以用minInt32 也可以用第一天的负值
		maxProfit := -prices[0]
		for j := 1; j < len(prices); j++ {
			maxProfit = max(maxProfit, dp[i-1][j-1]-prices[j-1])
			dp[i][j] = max(dp[i][j-1], prices[j]+maxProfit)
		}
	}
	return dp[k][len(prices)-1]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
这种dp更好理解
*/
class Solution {
    public int maxProfit(int k, int[] prices) {
        // write your code here
        if (k == 0 || prices.length == 0) {
            return 0;
        }
        if (k >= prices.length / 2) {
            int profit = 0;
            for (int i = 1; i < prices.length; i++) {
                if (prices[i] > prices[i - 1]) {
                    profit += prices[i] - prices[i - 1];
                }
            }
            return profit;
        }
        int n = prices.length;
        int[][] own = new int[n][k+1];   // own[i][j] 表示前i+1天，最多进行j次交易，第i天buy或者持有股票的最大收益
        int[][] sell = new int[n][k+1];  // sell[i][j] 表示前i+1天，进行j次交易，第i天sell或者不持有股票的最大获益

        for (int i = 0; i < n; i++) {
            for (int j = 1; j <= k; j++) {
                if(i == 0){
                    own[0][j] = - prices[i];
                }else{
                    own[i][j] = Math.max(own[i - 1][j], sell[i-1][j-1] - prices[i]);
                    sell[i][j] = Math.max(sell[i - 1][j], own[i-1][j] + prices[i]);
                }
            }
        }
        return sell[n - 1][k];
    }
}