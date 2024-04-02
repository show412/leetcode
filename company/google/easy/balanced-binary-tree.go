/*
 * @Author: hongwei.sun
 * @Date: 2021-01-22 18:45:52
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-02 09:43:48
 * @Description: file content
 */
import "math"

// https://leetcode.com/problems/balanced-binary-tree/
/*
Given a binary tree, determine if it is height-balanced.

For this problem, a height-balanced binary tree is defined as:

a binary tree in which the left and right subtrees of every node differ in height by no more than 1.



Example 1:

Given the following tree [3,9,20,null,null,15,7]:

    3
   / \
  9  20
    /  \
   15   7
Return true.

Example 2:

Given the following tree [1,2,2,3,3,null,null,4,4]:

       1
      / \
     2   2
    / \
   3   3
  / \
 4   4
Return false.
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	balanced, _ := checkDepth(root)
	return balanced
}

func checkDepth(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}
	leftBalanced, left := checkDepth(root.Left)
	
	rightBalanced, right := checkDepth(root.Right)
	balanced := (leftBalanced && rightBalanced && math.Abs(float64(left - right)) <=1) 
	return balanced, 1 + max(left, right)
}
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
