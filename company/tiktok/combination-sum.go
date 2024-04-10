/*
 * @Author: hongwei.sun
 * @Date: 2024-04-06 22:45:29
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-06 22:45:29
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

// 这个比较好理解，而且不用排序因为我不是通过candidate[i] < target 作为终止条件，是直接算和
var res [][]int
var entry []int
func combinationSum(candidates []int, target int) [][]int {
	// 注意这里要初始化，否则可能run多个case 全局变量还保持上次的数
    res = make([][]int, 0)
    entry = make([]int, 0)
	dfs(0, candidates, target)	
	return res
}

func dfs(start int, candidates []int, target int) {
	if sumWithForLoop(entry) == target {
		cpy := make([]int, len(entry))
		copy(cpy, entry)
		res = append(res, cpy)
	}
    if sumWithForLoop(entry) > target {
		return
	}
	for i := start; i < len(candidates); i++ {
		entry = append(entry, candidates[i])
		dfs(i, candidates, target)
		entry = entry[:len(entry)-1]
	}
}


func sumWithForLoop(numbers []int) int {
    sum := 0
    for _, num := range numbers {
        sum += num
    }
    return sum
}

// 这个的run time没有我上面写的好
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
