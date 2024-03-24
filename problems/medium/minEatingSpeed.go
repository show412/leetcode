/*
 * @Author: hongwei.sun
 * @Date: 2024-03-24 16:30:27
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-03-24 17:16:37
 * @Description: file content
 */
//  https://leetcode.com/problems/koko-eating-bananas/
// Koko loves to eat bananas. There are n piles of bananas, the ith pile has piles[i] bananas. The guards have gone and will come back in h hours.

// Koko can decide her bananas-per-hour eating speed of k. Each hour, she chooses some pile of bananas and eats k bananas from that pile. If the pile has less than k bananas, she eats all of them instead and will not eat any more bananas during this hour.

// Koko likes to eat slowly but still wants to finish eating all the bananas before the guards return.

// Return the minimum integer k such that she can eat all the bananas within h hours.

 

// Example 1:

// Input: piles = [3,6,7,11], h = 8
// Output: 4
// Example 2:

// Input: piles = [30,11,23,4,20], h = 5
// Output: 30
// Example 3:

// Input: piles = [30,11,23,4,20], h = 6
// Output: 23
 

// Constraints:

// 1 <= piles.length <= 104
// piles.length <= h <= 109
// 1 <= piles[i] <= 109
/*
1, it must be  h >= len(piles) otherwise it can't finish eat all
2, mostly, max piles in array can be one solution, but not best solution
so we need to use binary search to find proper value from 1 to max piles
TC is O(log max(piles))
*/
func minEatingSpeed(piles []int, h int) int {
    sort.Ints(piles)
	l, r := 1, piles[len(piles)-1]
	res := piles[len(piles)-1]
	for l <= r {
		k := l + (r-l)/2
		hours := 0
		for _, pile := range piles {
			// 必须得分别转float64, 否则int除以int 结果还是int 会向下进位
			hours += int(math.Ceil(float64(pile)/float64(k)))
		}
		// if hours == h {
		// 	res = k
		// 	break
		// } else if hours < h {
		// 	r = k-1
		// } else {
		// 	l = k+1
		// }
		/*
		because it doesn't have to be equal to h
		we just find minimum k
		*/
		if hours <= h {
			res = min(res, k)
			r = k-1
		} else {
			l = k+1
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}