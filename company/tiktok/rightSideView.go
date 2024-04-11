/*
 * @Author: hongwei.sun
 * @Date: 2024-04-11 16:16:59
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-11 16:16:59
 * @Description: file content
 */
// https://leetcode.com/problems/binary-tree-right-side-view/
/*
Given a binary tree, imagine yourself standing on the right side of it,
return the values of the nodes you can see ordered from top to bottom.

Example:

Input: [1,2,3,null,5,null,4]
Output: [1, 3, 4]
Explanation:

   1            <---
 /   \
2     3         <---
 \     \
  5     4       <---
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
binary tree BFS and get the last node as the right view
 case:
	[4,3,6,1,null,5,null,null,2]
	[]

*/
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := make([]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		size := len(queue)
		value := 0
		for i := 0; i < size; i++ {
			node := queue[0]
			if size > 1 {
				queue = queue[1:]
			} else {
				queue = []*TreeNode{}
			}
			value = node.Val
			if i == size-1 {
				res = append(res, value)
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return res
}
