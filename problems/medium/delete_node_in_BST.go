// https://leetcode.com/problems/delete-node-in-a-bst/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key {
		if root.Left == nil && root.Right == nil {
			root = nil
		} else if root.Left != nil && root.Right != nil {
			max := findMax(root.Left)
			root.Val = max
			root.Left = deleteNode(root.Left, max)
		} else if root.Left == nil {
			root = root.Right
		} else if root.Right == nil {
			root = root.Left
		}
	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else {
		root.Right = deleteNode(root.Right, key)
	}
	return root
}
func findMax(root *TreeNode) int {
	if root.Right == nil {
		return root.Val
	}
	return findMax(root.Right)
}
