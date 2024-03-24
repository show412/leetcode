/*
 * @Author: hongwei.sun
 * @Date: 2021-01-22 18:45:51
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-03-24 23:02:21
 * @Description: file content
 */
// https://leetcode.com/problems/diameter-of-binary-tree/
/*
Given a binary tree,
you need to compute the length of the diameter of the tree.
The diameter of a binary tree is the length of the longest path
between any two nodes in a tree.
This path may or may not pass through the root.

Example:
Given a binary tree
          1
         / \
        2   3
       / \
      4   5
Return 3, which is the length of the path [4,2,1,3] or [5,2,1,3].

Note: The length of path between two nodes is represented
by the number of edges between them.
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 /*
 计算通过每个node的diameter，用递归
 从root出发 其实就是算左右子树种最深的情况下子节点的总和
 */
//  define one global diameter we can access it in each function
var diameter int
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	diameter = 0
	depth(root)
	return diameter
}
func depth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	l := depth(node.Left)
	r := depth(node.Right)
	// this is key, l+r is current diameter through node
	diameter = max(diameter, l+r)
	// return the edge length for parent node, notice need to plus 1 
	// because we calculate current node
	return max(l, r) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}