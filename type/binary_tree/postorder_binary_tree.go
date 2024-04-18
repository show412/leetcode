/*
 * @Author: hongwei.sun
 * @Date: 2024-04-18 17:46:28
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-18 17:50:35
 * @Description: file content
 */
//  https://leetcode.com/problems/binary-tree-postorder-traversal/description/
/*
Given the root of a binary tree, return the postorder traversal of its nodes' values.

 

Example 1:


Input: root = [1,null,2,3]
Output: [3,2,1]
Example 2:

Input: root = []
Output: []
Example 3:

Input: root = [1]
Output: [1]
 

Constraints:

The number of the nodes in the tree is in the range [0, 100].
-100 <= Node.val <= 100
 

Follow up: Recursive solution is trivial, could you do it iteratively?
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func postorderTraversal(root *TreeNode) []int {
    res := make([]int, 0)
	if root == nil {
		return res
	}
	if root.Left != nil {
		res = append(res, postorderTraversal(root.Left)...)
	}
	if root.Right != nil {
		res = append(res, postorderTraversal(root.Right)...)
	}
	res = append(res, root.Val)
	return res
 }