/*
 * @Author: hongwei.sun
 * @Date: 2021-01-22 18:45:52
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-03-22 12:22:39
 * @Description: file content
 */
// https://leetcode.com/problems/permutations/
/*
Given an array nums of distinct integers, return all the possible permutations. You can return the answer in any order.
Example 1:

Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
Example 2:

Input: nums = [0,1]
Output: [[0,1],[1,0]]
Example 3:

Input: nums = [1]
Output: [[1]]
 
Constraints:

1 <= nums.length <= 6
-10 <= nums[i] <= 10
All the integers of nums are unique.
*/

/*
permutation 所有的排列组合
[1,2,3]  可以选1或者2或者3 然后往下类推
要在算法中实现的是 如何避免重复选和回溯时候去掉当前选的num
    1      2    3    
   / \    /\   / \
  2  3   1 3  1  2 
 /  /   /  \  \  \
3  2   3   2  3  1
*/
func permute(nums []int) [][]int {
	res := [][]int{}
	// 避免重复 用hash
	visit := make(map[int]bool)
	
	entry := []int{}
	dfs(nums, visit, &entry, &res)
	return res
}

func dfs(nums []int, visit map[int]bool, entry *[]int, res *[][]int) {
	if len(*entry) == len(nums) {
		cpy := make([]int, len(*entry))
		copy(cpy, *entry)
		*res = append(*res, cpy)
		return
	}

	for i := 0; i < len(nums); i++ {

		if visit[nums[i]] == true {
			continue
		} else {
			visit[nums[i]] = true
			*entry = append(*entry, nums[i])
			dfs(nums, visit, entry, res)
			// 回溯时候去掉当前选的num
			visit[nums[i]] = false
			*entry = (*entry)[:len(*entry)-1]
		}
	}
	return
}
