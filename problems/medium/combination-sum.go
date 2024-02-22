/*
 * @Author: hongwei.sun
 * @Date: 2021-01-22 18:45:52
 * @LastEditors: your name
 * @LastEditTime: 2024-02-20 17:57:59
 * @Description: file content
 */
import "sort"

// https://leetcode.com/problems/combination-sum/
/*
Given a set of candidate numbers (candidates) (without duplicates) and a target number (target), find all unique combinations in candidates where the candidate numbers sums to target.

The same repeated number may be chosen from candidates unlimited number of times.

Note:

All numbers (including target) will be positive integers.
The solution set must not contain duplicate combinations.
Example 1:

Input: candidates = [2,3,6,7], target = 7,
A solution set is:
[
  [7],
  [2,2,3]
]
Example 2:

Input: candidates = [2,3,5], target = 8,
A solution set is:
[
  [2,2,2,2],
  [2,3,3],
  [3,5]
]
*/
func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	if len(candidates) == 0 {
		return res
	}
	// we need to sort this array otherwise it will miss answer.
	// cause we use target-candidate to decide when to break
	sort.Ints(candidates)
	start := 0
	combination := make([]int, 0)
	dfs(candidates, start, &combination, target, &res)
	return res
}

func dfs(candidates []int, start int, combination *[]int, target int, res *[][]int) {
	// use traget to minus, can avoid to bring another variable total as paramters in function
	if target == 0 {
		cpy := make([]int, len(*combination))
		copy(cpy, *combination)
		*res = append(*res, cpy)
		return
	}
	for i := start; i < len(candidates); i++ {
		if target < candidates[i] {
			break
		}
		*combination = append(*combination, candidates[i])
		// transmit i not i+1 means the i number could be repeated
		// this is key
		dfs(candidates, i, combination, target-candidates[i], res)
		// notice the pointer needs bracket to be priority
		*combination = (*combination)[:len(*combination)-1]
	}
}
