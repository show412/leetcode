/*
 * @Author: hongwei.sun
 * @Date: 2021-01-22 18:45:51
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-03-30 23:14:07
 * @Description: file content
 */
// https://leetcode.com/problems/binary-tree-inorder-traversal/
/*
Given a binary tree, return the inorder traversal of its nodes' values.

Example:

Input: [1,null,2,3]
   1
    \
     2
    /
   3

Output: [1,3,2]
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
//  there are two other solution without recursion
// https://leetcode.com/problems/binary-tree-inorder-traversal/solution/
func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	left := inorderTraversal(root.Left)
	right := inorderTraversal(root.Right)
	if len(left) != 0 {
		// ...代表可变参数，slice的话就是把这个slice里的value都append到res里
		res = append(res, left...)
	}
	res = append(res, root.Val)
	if len(right) != 0 {
		res = append(res, right...)
	}
	return res
}
