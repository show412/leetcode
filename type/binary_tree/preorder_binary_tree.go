/*
 * @Author: hongwei.sun
 * @Date: 2024-04-18 17:46:09
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-18 17:46:09
 * @Description: file content
 */
// https://leetcode.com/problems/binary-tree-preorder-traversal/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func preorderTraversal(root *TreeNode) []int {
  result := []int
  result = traverse(root, result)
  return result
}

func traverse(root *TreeNode, result []int) []int {
  if root == nil {
    return result
  }
  result = append(result, root.Val)
  result = traverse(root.Left, result)
  result = traverse(root.Right, result)
  return result
}