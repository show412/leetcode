/*
 * @Author: hongwei.sun
 * @Date: 2021-01-22 18:45:52
 * @LastEditors: your name
 * @LastEditTime: 2024-02-15 17:29:12
 * @Description: file content
 */
// https://leetcode.com/problems/maximum-depth-of-binary-tree/
/*
Given a binary tree, find its maximum depth.

The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

Note: A leaf is a node with no children.

Example:

Given binary tree [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
return its depth = 3.
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	depth := 0
	if root == nil {
		return depth
	} else {
		depth++
		depth += max(maxDepth(root.Left), maxDepth(root.Right))
	}
	return depth
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
