/*
 * @Author: hongwei.sun
 * @Date: 2024-04-03 14:54:08
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-16 23:17:38
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
 
 func spiralOrder(matrix [][]int) []int {
    
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
