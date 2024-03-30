/*
 * @Author: hongwei.sun
 * @Date: 2024-03-30 23:18:11
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-03-30 23:39:54
 * @Description: file content
 */
// https://leetcode.com/problems/binary-tree-level-order-traversal/description/
// Given the root of a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level).

 

// Example 1:


// Input: root = [3,9,20,null,null,15,7]
// Output: [[3],[9,20],[15,7]]
// Example 2:

// Input: root = [1]
// Output: [[1]]
// Example 3:

// Input: root = []
// Output: []
 

// Constraints:

// The number of nodes in the tree is in the range [0, 2000].
// -1000 <= Node.val <= 1000
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//  BFS solution
func levelOrder(root *TreeNode) [][]int {
    res := make([][]int,0)
	if root == nil {
		return res
	}
	q := make([]*TreeNode,0)
	q = append(q, root)
	for len(q) != 0 {
		qlen := len(q)
		level := make([]int,0)
		// 这要注意是要拿当前queue的长度 要提前算出来 不能动态的len(q)
		// 否则的话 下面的for循环会继续走下去，把下一层的值append到上一层
		for i := 0; i < qlen; i++ {
			// pop first every time
			node := q[0]
			q = q[1:]
			level = append(level, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		res = append(res, level)
	}
	return res
 }

