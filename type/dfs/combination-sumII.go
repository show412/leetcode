/*
 * @Author: hongwei.sun
 * @Date: 2024-04-06 22:48:37
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-06 22:57:48
 * @Description: file content
 */
//  https://leetcode.com/problems/combination-sum-ii/description/
/*
Given a collection of candidate numbers (candidates) and a target number (target), find all unique combinations in candidates where the candidate numbers sum to target.

Each number in candidates may only be used once in the combination.

Note: The solution set must not contain duplicate combinations.

 

Example 1:

Input: candidates = [10,1,2,7,6,1,5], target = 8
Output: 
[
[1,1,6],
[1,2,5],
[1,7],
[2,6]
]
Example 2:

Input: candidates = [2,5,2,1,2], target = 5
Output: 
[
[1,2,2],
[5]
]
 

Constraints:

1 <= candidates.length <= 100
1 <= candidates[i] <= 50
1 <= target <= 30
*/

var res [][]int
var entry []int
func combinationSum2(candidates []int, target int) [][]int {
	// 注意这里要初始化，否则可能run多个case 全局变量还保持上次的数
    res = make([][]int, 0)
    entry = make([]int, 0)
    sort.Ints(candidates)
	dfs(0, candidates, target)	
	return res
}

func dfs(start int, candidates []int, target int) {
	if sumWithForLoop(entry) == target {
		cpy := make([]int, len(entry))
		copy(cpy, entry)
		res = append(res, cpy)
		return
	}
    if sumWithForLoop(entry) > target {
		return
	}
	for i := start; i < len(candidates); i++ {
        entry = append(entry, candidates[i])
        dfs(i+1, candidates, target)
        entry = entry[:len(entry)-1]
        for i+1 < len(candidates) && candidates[i+1] == candidates[i] {
            i++
        }
	}
}

func sumWithForLoop(numbers []int) int {
    sum := 0
    for _, num := range numbers {
        sum += num
    }
    return sum
}