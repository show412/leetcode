import "math"

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/
/*
You are given an array prices where prices[i] is the price of a given stock on the ith day.

Find the maximum profit you can achieve. You may complete at most two transactions.

Note: You may not engage in multiple transactions simultaneously 
(i.e., you must sell the stock before you buy again).

 

Example 1:

Input: prices = [3,3,5,0,0,3,1,4]
Output: 6
Explanation: Buy on day 4 (price = 0) and sell on day 6 (price = 3), profit = 3-0 = 3.
Then buy on day 7 (price = 1) and sell on day 8 (price = 4), profit = 4-1 = 3.
Example 2:

Input: prices = [1,2,3,4,5]
Output: 4
Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
Note that you cannot buy on day 1, buy on day 2 and sell them later, as you are engaging multiple transactions at the same time. You must sell before buying again.
Example 3:

Input: prices = [7,6,4,3,1]
Output: 0
Explanation: In this case, no transaction is done, i.e. max profit = 0.
 

Constraints:

1 <= prices.length <= 105
0 <= prices[i] <= 105
*/
// 两次交易用左右两个滑动窗口标识
// 两个数组left right表示左右两个在相应位置的max profit
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	res := 0
	// 从左开始的最大profit
	left := make([]int, len(prices))
	// 从右开始的最大profit
  	right := make([]int, len(prices))
	// 算左边在每个i的最大profit
	maxProfitLeft := 0
	minPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		minPrice = min(prices[i], minPrice)
		maxProfitLeft = max(maxProfitLeft, prices[i]-minPrice)
		left[i] = maxProfitLeft
	}
	// 算右边在每个i的最大profit
	maxProfitRight := 0
	maxPrice := prices[len(prices)-1]
	for i := len(prices) - 1; i >= 0; i-- {
		maxPrice = max(prices[i], maxPrice)
		maxProfitRight = max(maxProfitRight, maxPrice-prices[i])
		right[i] = maxProfitRight
	}
	// 找左右两边一起相加 或者只有left(一次交易)的max profit
	for i := 0; i < len(left); i++ {
		profit := left[i]
		if i != len(left) - 1 {
			profit += right[i+1]
		}
		res = max(res, profit)
	}
	return res
}

 func max(a, b int) int {
	if a > b {
		return a
	}
	return b
 }

 func min(a, b int) int {
	if a < b {
		return a
	}
	return b
 }

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
 }

 func min(a, b int) int {
	if a < b {
		return a
	}
	return b
 }

// straight forward solution O(n^2)
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	res := 0
	l1 := 0
	r1 := 1
	for l1 < r1 && l1 < len(prices) {
        r1 = l1 + 1
		for r1 < len(prices) {
			buy1 := prices[l1]
			sell1 := prices[r1]
			profit1 := max(sell1 - buy1, 0)
            if r1 < len(prices)-2 {
                l2 := r1 + 1
			    r2 := l2 + 1
			    profit2 := 0
                for l2 < r2 && l2 < len(prices) {
                    r2 = l2 + 1
                    for r2 < len(prices) {
                        buy2 := prices[l2]
                        sell2 := prices[r2]
                        profit2 = max(sell2 - buy2, 0)
                        res = max(res, profit1+profit2)
                        r2++
                    }
				l2++
			    }
            } 
            res = max(res, profit1)
			r1++
		}
		l1++
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}


// a smart solution
// 就是按正常思维去解 先买再卖 第二次买再卖 就是不好想到

/*
看算法感觉好像是都在一个点买卖 其实不是
hold1是为了release1 release1为了hold2 hold2为了release2
其实hold1 release1 hold2 release2不可能在同一点
并且hold1 < release1 < hold2 < release2
通过  = max() 不停地算 最后得到release2
其实hold1 release1 hold2的价格都是不对的 不是实际能得到结果时的price
*/
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	hold1 := math.MinInt32
	hold2 := math.MinInt32
	release1 := 0
	release2 := 0
	for i := 0; i < len(prices); i++ {
		price := prices[i]
		hold1 = max(hold1, -price)
		release1 = max(release1, hold1+price)
		hold2 = max(hold2, release1-price)
		release2 = max(release2, hold2+price)
	}
	return release2
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// DP, refer to the thought with cooldown
