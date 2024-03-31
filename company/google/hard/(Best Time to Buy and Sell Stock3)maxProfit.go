import "math"

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/
// twice almost
// there is line to divide the array into two
// and we caculate the two array maxProfit respectively
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	res := 0
	for i := 0; i < len(prices); i++ {
		first := make([]int, len(prices[:i+1]))
		second := make([]int, len(prices[i+1:]))
		copy(first, prices[:i+1])
		copy(second, prices[i+1:])
		firstProfit := 0
		firstMin := math.MaxInt32
		secondProfit := 0
		secondMin := math.MaxInt32

		for j := 0; j < len(first); j++ {
			if first[j] < firstMin {
				firstMin = first[j]
			}
			if firstProfit < (first[j] - firstMin) {
				firstProfit = first[j] - firstMin
			}
		}

		for j := 0; j < len(second); j++ {
			if second[j] < secondMin {
				secondMin = second[j]
			}
			if secondProfit < (second[j] - secondMin) {
				secondProfit = second[j] - secondMin
			}
		}
		if res < (firstProfit + secondProfit) {
			res = firstProfit + secondProfit
		}
	}
	return res
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
