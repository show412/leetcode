/*
 * @Author: hongwei.sun
 * @Date: 2023-10-11 10:48:14
 * @LastEditors: your name
 * @LastEditTime: 2023-10-18 16:32:07
 * @Description: file content
https://leetcode.com/problems/same-tree/
Given the roots of two binary trees p and q, write a function to check if they are the same or not.
Two binary trees are considered the same
if they are structurally identical, and the nodes have the same value.
it's about recursion solution 递归
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
