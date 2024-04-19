/*
 * @Author: hongwei.sun
 * @Date: 2024-04-03 14:54:08
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-19 11:46:04
 * @Description: file content
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 func canJump(nums []int) bool {
	goal := len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] + i >= goal {
			goal = i
		}
	}
	if goal == 0 {
		return true
	}
	return false
 }


 var res [][]int
 var visit map[int]bool
 var candidate []int
 
 var res [][]int
 var candidate []int
 
func subsets(nums []int) [][]int {
	res = make([][]int, 0)
	candidate = make([]int, 0)
	dfs(nums, 0)
	return res
}

func dfs(nums []int, start int) {
	cpy := make([]int, len(candidate))
	copy(cpy, candidate)
	res = append(res, cpy)
	for i := start; i < len(nums); i++ {
		candidate = append(candidate, nums[i])
		dfs(nums, i+1)
		candidate = candidate[:len(candidate)-1]
	}
}

 
 
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
 
 func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	for root != nil {
		if root.Val < p.Val && root.Val < q.Val {
			root = root.Right
		} else if root.Val > p.Val && root.Val > q.Val {
			root = root.Left
		} else {
			break
		}
	}
	return root
 }

 func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if inNode(p, q) {
		return p
	}
	if inNode(q, p) {
		return q
	}
    var res *TreeNode
	for root != nil {
		if inNode(root.Left, p) && inNode(root.Left, q) {
			root = root.Left
		} else if inNode(root.Right, p) && inNode(root.Right, q) {
			root = root.Right
		} else {
			res = root
            break
		}
	}
    return res
 }

 func inNode(parent *TreeNode, child *TreeNode) bool {
	if parent == child {
		return true
	}
	if parent == nil {
		return false
	}
	return inNode(parent.Left, child) || inNode(parent.Right, child)
 }
 
func sumWithForLoop(numbers []int) int {
    sum := 0
    for _, num := range numbers {
        sum += num
    }
    return sum
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
