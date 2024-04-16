/*
 * @Author: hongwei.sun
 * @Date: 2024-04-16 23:05:17
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-16 23:05:18
 * @Description: file content
 */
// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/
/*
Given a binary tree, find the lowest common ancestor (LCA) of
two given nodes in the tree.

According to the definition of LCA on Wikipedia:
“The lowest common ancestor is defined between two nodes p and q
as the lowest node in T that has both p and q as descendants
(where we allow a node to be a descendant of itself).”

Given the following binary tree:  root = [3,5,1,6,2,0,8,null,null,7,4]




Example 1:

Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
Output: 3
Explanation: The LCA of nodes 5 and 1 is 3.
Example 2:

Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
Output: 5
Explanation: The LCA of nodes 5 and 4 is 5,
since a node can be a descendant of itself according to the LCA definition.


Note:

All of the nodes' values will be unique.
p and q are different and both values will exist in the binary tree.
*/
/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
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
