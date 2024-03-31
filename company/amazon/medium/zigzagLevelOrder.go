// https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/
/*
Given a binary tree, return the zigzag level order traversal of its nodes' values.
(ie, from left to right, then right to left for the next level and alternate between).

For example:
Given binary tree [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
return its zigzag level order traversal as:
[
  [3],
  [20,9],
  [15,7]
]
*/
/**
[3,9,20,null,null,15,7]
[
  [3],
  [20,9],
  [15,7]
]

[1,2,3,4,null,null,5]
[[1],[3,2],[4,5]]

 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
*/
func zigzagLevelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	k := 1
	for len(queue) != 0 {
		size := len(queue)
		level := make([]int, 0)
		for i := 0; i < size; i++ {
			node := queue[0]
			if size > 1 {
				queue = queue[1:]
			} else {
				queue = []*TreeNode{}
			}
			value := node.Val
			if k%2 == 1 {
				level = append(level, value)
			} else {
				// this is key, make value into first index
				level = append([]int{value}, level...)
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, level)
		k++
	}
	return res
}