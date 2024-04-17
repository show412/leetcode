/*
 * @Author: hongwei.sun
 * @Date: 2024-04-17 10:43:51
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-17 12:30:08
 * @Description: file content
 */
// https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/
/*
Given two integer arrays preorder and inorder where preorder is the preorder traversal of a binary tree and inorder is the inorder traversal of the same tree, construct and return the binary tree.

 

Example 1:


Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
Output: [3,9,20,null,null,15,7]
Example 2:

Input: preorder = [-1], inorder = [-1]
Output: [-1]
 

Constraints:

1 <= preorder.length <= 3000
inorder.length == preorder.length
-3000 <= preorder[i], inorder[i] <= 3000
preorder and inorder consist of unique values.
Each value of inorder also appears in preorder.
preorder is guaranteed to be the preorder traversal of the tree.
inorder is guaranteed to be the inorder traversal of the tree.
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
1, 从preorder能知道root
2，从inorder能知道root的位置，进而知道左子树有多少个节点，右子树有多少个节点
3，然后分别递归得到左右子树
*/
 func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{Val:preorder[0]}
	var rootIndex int
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			rootIndex = i
			break
		}
	}
	/*
	by rootIndex we can know how many left nodes and how many right nodes
	extract left node numbers and right node numbers from preorder and inorder repectively
	then recusion to build tree from left and right tree
	rootIndex also means how many nodes is on left of root in inorder array
	so we extract such numbers from preorder to preorder as left sub tree
	*/
	root.Left = buildTree(preorder[1:rootIndex+1], inorder[:rootIndex])
	root.Right = buildTree(preorder[rootIndex+1:], inorder[rootIndex+1:])
	return root
 }