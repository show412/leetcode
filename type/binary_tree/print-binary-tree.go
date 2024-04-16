/*
 * @Author: hongwei.sun
 * @Date: 2024-04-16 13:11:46
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-16 13:11:46
 * @Description: file content
 */
//  https://leetcode.com/problems/print-binary-tree/description/
/*
Given the root of a binary tree, construct a 0-indexed m x n string matrix res that represents a formatted layout of the tree. The formatted layout matrix should be constructed using the following rules:

The height of the tree is height and the number of rows m should be equal to height + 1.
The number of columns n should be equal to 2height+1 - 1.
Place the root node in the middle of the top row (more formally, at location res[0][(n-1)/2]).
For each node that has been placed in the matrix at position res[r][c], place its left child at res[r+1][c-2height-r-1] and its right child at res[r+1][c+2height-r-1].
Continue this process until all the nodes in the tree have been placed.
Any empty cells should contain the empty string "".
Return the constructed matrix res.

 

Example 1:


Input: root = [1,2]
Output: 
[["","1",""],
 ["2","",""]]
Example 2:


Input: root = [1,2,3,null,4]
Output: 
[["","","","1","","",""],
 ["","2","","","","3",""],
 ["","","4","","","",""]]
 

Constraints:

The number of nodes in the tree is in the range [1, 210].
-99 <= Node.val <= 99
The depth of the tree will be in the range [1, 10].
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 var res [][]string
 var height int
 func printTree(root *TreeNode) [][]string {
    height = getHeight(root) - 1
	columns := int(math.Pow(2, float64(height+1))-1)
	res = make([][]string, height+1)
	for i := 0; i < height+1; i++ {
		res[i] = make([]string, columns)
	}
	traverse(root, 0, (columns-1)/2)
	return res
 }

 func traverse(node *TreeNode, r int, c int) {
	if node == nil {
		return
	}
	res[r][c] = strconv.Itoa(node.Val)
	lc, rc :=  c - int(math.Pow(2, float64(height - r - 1))), c + int(math.Pow(2, float64(height - r - 1)))
	traverse(node.Left, r+1, lc)
	traverse(node.Right, r+1, rc)
 }


 func getHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return max(getHeight(node.Left), getHeight(node.Right)) + 1
 }