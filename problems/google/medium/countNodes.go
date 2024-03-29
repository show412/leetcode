// https://leetcode.com/problems/count-complete-tree-nodes/
/*
Given a complete binary tree, count the number of nodes.

Note:

Definition of a complete binary tree from Wikipedia:
In a complete binary tree every level,
except possibly the last, is completely filled,
and all nodes in the last level are as far left as possible.
It can have between 1 and 2h nodes inclusive at the last level h.

Example:

Input:
    1
   / \
  2   3
 / \  /
4  5 6

Output: 6
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//  TC is O(N) SC is O(h), h is the high of treeNode
// treeNode almost could use divide conquer to solve
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	nums := 1
	if root.Left != nil {
		nums += countNodes(root.Left)
	}
	if root.Right != nil {
		nums += countNodes(root.Right)
	}
	return nums
}
