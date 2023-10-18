/*
 * @Author: hongwei.sun
 * @Date: 2023-10-18 16:25:44
 * @LastEditors: your name
 * @LastEditTime: 2023-10-18 16:35:45
 * @Description: file content
 https://leetcode.com/problems/symmetric-tree/
 Given the root of a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSym(root.Left, root.Right)
}

func isSym(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil || left.Val != right.Val {
		return false
	}
	//should be mirror return isSym(left.Left, left.Right) && isSym(right.Left, right.Right)
	return isSym(left.Left, right.Right) && isSym(left.Right, right.Left)
}