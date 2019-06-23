import "math"

// https://leetcode.com/problems/perfect-squares/
/*
Given a positive integer n, find the least number of perfect square numbers
(for example, 1, 4, 9, 16, ...) which sum to n.

Example 1:

Input: n = 12
Output: 3
Explanation: 12 = 4 + 4 + 4.
Example 2:

Input: n = 13
Output: 2
Explanation: 13 = 4 + 9.
*/
// LTE need to improve for 7168
func numSquares(n int) int {
	if n == 1 || n == 4 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if n == 3 {
		return 3
	}
	var list []int
	for i := 1; i < n; i++ {
		if i*i > n {
			break
		}
		list = append(list, i)
	}
	if list[len(list)-1]*list[len(list)-1] == n {
		return 1
	}
	var sub []int
	cnt := 0
	var res int
	dfs(list, 0, sub, cnt, &res, n)

	return res
}

func dfs(list []int, start int, sub []int, cnt int, res *int, max int) {

	for i := start; i < len(list); i++ {
		sub = append(sub, list[i])
		cnt += list[i] * list[i]
		if cnt == max {
			if *res == 0 || *res > len(sub) {
				*res = len(sub)
				sub = sub[:len(sub)-1]
				cnt -= list[i] * list[i]
				break
			}
		}
		if *res != 0 && *res < len(sub) {
			sub = sub[:len(sub)-1]
			cnt -= list[i] * list[i]
			break
		}
		dfs(list, i, sub, cnt, res, max)
		sub = sub[:len(sub)-1]
		cnt -= list[i] * list[i]
	}

}

// good performance with DFS
func numSquares(n int) int {
	cache := make(map[int]int)
	cache[0] = 0
	return doNumSquares(n, cache)
}

func doNumSquares(n int, cache map[int]int) int {
	if found, ok := cache[n]; ok {
		return found
	}

	res := math.MaxInt32
	for i := 1; i*i <= n; i++ {
		target := n - i*i
		res = min(res, doNumSquares(target, cache)+1)
	}
	cache[n] = res
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
