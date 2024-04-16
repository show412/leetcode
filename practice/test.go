/*
 * @Author: hongwei.sun
 * @Date: 2024-04-03 14:54:08
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-16 11:43:40
 * @Description: file content
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	q := make([]*TreeNode, 0)
	q = append(q, root)
	for len(q) > 0 && q[0] != nil {
		node := q[0]
		q = append(q, node.Left, node.Right)
		q = q[1:]
	}
	for i := 0; i < len(q); i++ {
		if q[i] != nil {
			return false
		}
	}

	return true
 }

 func traverse
 


func sumWithForLoop(numbers []int) int {
    sum := 0
    for _, num := range numbers {
        sum += num
    }
    return sum
}

 func max(a, b int) int {
	if a > b {
		return a
	}
	return b
 }

 func min(a, b int) int {
	if a < b {
		return a
	}
	return b
 }
