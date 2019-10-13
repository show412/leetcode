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
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	diameter := 1
	_ = depth(root, &diameter)
	return diameter - 1
}
func depth(node *TreeNode, diameter *int) int {
	if node == nil {
		return 0
	}
	l := depth(node.Left, diameter)
	r := depth(node.Right, diameter)
	// this is key, 总会有一个节点的左右深度为最大
	*diameter = max(*diameter, l+r+1)
	return max(l, r) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}