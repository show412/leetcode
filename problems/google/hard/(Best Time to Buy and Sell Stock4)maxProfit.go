// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv/
/*
Say you have an array for which the ith element is the price of a given stock on day i.

Design an algorithm to find the maximum profit. You may complete at most k transactions.

Note:
You may not engage in multiple transactions at the same time (ie, you must sell the stock before you buy again).

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
	profit := 0
	if k >= len(prices)/2 {
		for i := 1; i < len(prices); i++ {
			if prices[i] > prices[i-1] {
				profit += (prices[i] - prices[i-1])
			}
		}
		return profit
	}
	// dp[i][j] means the max profit j days finish i times transaction
	dp := make([][]int, k+1)
	for i := 0; i <= k; i++ {
		dp[i] = make([]int, len(prices))
	}
	dp[0][0] = 0
	for i := 0; i < len(dp); i++ {
		dp[i][0] = 0
	}
	for j := 0; j < len(dp[0]); j++ {
		dp[0][j] = 0
	}
	// function
	for i := 1; i <= k; i++ {
		buy := -prices[0]
		for j := 1; j < len(prices); j++ {
			//
			dp[i][j] = max(dp[i][j-1], prices[j]+buy)
			buy = max(buy, dp[i-1][j-1]-prices[j])
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




// seems it doesn't work
// func maxProfit(k int, prices []int) int {
//   if len(prices) <=1 || k==0 {
// 		return 0
// 	}
// 	profit := 0
// 	if k >= len(prices)/2 {
// 		for i :=1; i < len(prices); i++ {
// 			if prices[i] > prices[i-1] {
// 				profit += (prices[i] - prices[i-1])
// 			}
// 		}
// 		return profit
// 	}
// 	// globalBest[i][j] means the best profit at the before i day with j times transactions
// 	// but the i day DOESN'T must to sell
// 	globalBest := make([][]int, len(prices))
// 	for i :=0; i < len(globalBest); i++ {
// 		globalBest[i] := make([]int, len(prices))
// 	}
// 	// define the state
// 	// mustSell[i][j] means the best profit at the before i day with j times transaction
// 	// and the i day must to sell
// 	mustSell := make([][]int, len(prices))
// 	for i :=0; i < len(mustSell); i++ {
// 		mustSell[i] := make([]int, len(prices))
// 	}
// 	// init
// 	globalBest[0][0] = 0
// 	mustSell[0][0] = 0
// 	for i :=0; i < k; i++ {
// 		globalBest[0][i] = 0
// 		mustSell[0][i] = 0
// 	}
// 	// function
// 	for i := 1; i<len(prices); i++ {
// 		price := prices[i] - price[i-1]
// 		globalBest[i][0] = 0
// 		mustSell[i][0] = 0
// 		for j :=1; j<k; j++ {
// 			globalBest[i][j] = max(globalBest[i-1][j], mustSell[i][j])
// 			mustSell[i][j] = max(globalBest[i-1][j-1]+price, mustSell[i-1][j]+price)
// 		}
// 	}
// 	return globalBest[len(prices)-1][k]
// }

