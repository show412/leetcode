import "math"

// https://leetcode.com/problems/closest-binary-search-tree-value/
/*
Given a non-empty binary search tree and a target value,
find the value in the BST that is closest to the target.

Note:

Given target value is a floating point.
You are guaranteed to have only one unique value in the BST
that is closest to the target.
Example:

Input: root = [4,2,5,1,3], target = 3.714286

    4
   / \
  2   5
 / \
1   3

Output: 4
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//  TC is O(h), h is the height of binary tree and SC is O(1)
// binary seach for bianry tree
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func closestValue(root *TreeNode, target float64) int {
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	closest := root.Val
	for root != nil {
		if target > float64(root.Val) {
			root = root.Right
		} else {
			root = root.Left
		}
		if root != nil && math.Abs(float64(root.Val)-target) < math.Abs(float64(closest)-target) {
			closest = root.Val
		}
	}
	return closest
}
