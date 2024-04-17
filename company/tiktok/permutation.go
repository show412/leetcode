/*
 * @Author: hongwei.sun
 * @Date: 2024-04-17 22:04:22
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-17 22:04:23
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

var res [][]int
var visit map[int]bool
var candidate []int
func permute(nums []int) [][]int {
   res = make([][]int, 0)
   visit = make(map[int]bool, len(nums))
   candidate = make([]int, 0)
   dfs(nums)
   return res
}

func dfs(nums []int) {
   if len(candidate) == len(nums) {
	   cpy := make([]int, len(candidate))
	   copy(cpy, candidate)
	   res = append(res, cpy)
	   return
   }
   for i := 0; i < len(nums); i++ {
	   if visit[nums[i]] == false {
		   candidate = append(candidate, nums[i])
		   visit[nums[i]] = true
		   dfs(nums)
		   candidate = candidate[:len(candidate)-1]
		   visit[nums[i]] = false
	   }
   }
}


// another writing
func permute(nums []int) [][]int {
	res := [][]int{}
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
			visit[nums[i]] = false
			*entry = (*entry)[:len(*entry)-1]
		}
	}
	return
}
