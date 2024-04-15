/*
 * @Author: hongwei.sun
 * @Date: 2024-04-14 22:56:45
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-15 16:52:43
 * @Description: file content
 */
// https://leetcode.com/problems/path-sum-ii/description/
/*
Given the root of a binary tree and an integer targetSum, return all root-to-leaf paths where the sum of the node values in the path equals targetSum. Each path should be returned as a list of the node values, not node references.

A root-to-leaf path is a path starting from the root and ending at any leaf node. A leaf is a node with no children.

 

Example 1:


Input: root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
Output: [[5,4,11,2],[5,8,4,5]]
Explanation: There are two paths whose sum equals targetSum:
5 + 4 + 11 + 2 = 22
5 + 8 + 4 + 5 = 22
Example 2:


Input: root = [1,2,3], targetSum = 5
Output: []
Example 3:

Input: root = [1,2], targetSum = 0
Output: []
 

Constraints:

The number of nodes in the tree is in the range [0, 5000].
-1000 <= Node.val <= 1000
-1000 <= targetSum <= 1000
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 var res [][]int
 var candidate []int
 func pathSum(root *TreeNode, targetSum int) [][]int {
    res = make([][]int, 0)
	candidate = make([]int, 0)
    if root == nil {
        return res
    }
    candidate = append(candidate, root.Val)
	traverse(root, targetSum - root.Val)
	return res
 }

 func traverse(node *TreeNode, targetSum int) {
	if node.Left == nil && node.Right == nil && targetSum == 0 {
		cpy := make([]int, len(candidate))
		copy(cpy, candidate)
		res = append(res, cpy)
		return
	}
	if node.Left != nil {
        candidate = append(candidate, node.Left.Val)
	    traverse(node.Left, targetSum - node.Left.Val)
	    candidate = candidate[:len(candidate)-1]
    }
	if node.Right != nil {
        candidate = append(candidate, node.Right.Val)
	    traverse(node.Right, targetSum - node.Right.Val)
	    candidate = candidate[:len(candidate)-1]
    }
 }