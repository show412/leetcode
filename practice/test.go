/*
 * @Author: hongwei.sun
 * @Date: 2024-04-03 14:54:08
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-09 20:50:31
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
 func isSymmetric(root *TreeNode) bool {
    if root == nil {
		return true
	}
	left := root.Left
	right := root.Right
	
	return isMirror(left,right)
 }

 func isMirror(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left != nil && right != nil && left.Val == right.Val {
		return isMirror(left.Left, right.Right) && isMirror(left.Right, right.Left)
	} else {
		return false
	}
 }




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
