import (
	"math"
	"sort"
)

// https://leetcode.com/problems/closest-binary-search-tree-value/
/*
Given a non-empty binary search tree and a target value,
find the value in the BST that is closest to the target.

Note:

Given target value is a floating point.
You are guaranteed to have only one unique value in the BST
that is closest to the target.
Example:

Input: root = [4,2,5,1,3], target = 3.714286

    4
   / \
  2   5
 / \
1   3

Output: 4
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//  TC and SC both are O(n), actually we could use binary search directly
func closestValue(root *TreeNode, target float64) int {
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	inOrderArray := make([]int, 0)
	inOrder(root, &inOrderArray)
	i := sort.Search(len(inOrderArray), func(i int) bool { return float64(inOrderArray[i]) >= target })
	if i == len(inOrderArray) || (i-1 >= 0 && (math.Abs(target-float64(inOrderArray[i]))) > (math.Abs(float64(inOrderArray[i-1])-target))) {
		return inOrderArray[i-1]
	} else {

		return inOrderArray[i]
	}
}

func inOrder(node *TreeNode, array *[]int) {
	if node.Left != nil {
		inOrder(node.Left, array)
	}
	*array = append(*array, node.Val)
	if node.Right != nil {
		inOrder(node.Right, array)
	}
}

// binary seach for bianry tree
// TC is O(h) h is the height of the tree,
func closestValue(root *TreeNode, target float64) int {
	value, res := root.Val, root.Val
	for root != nil {
		value = root.Val
		if math.Abs(float64(value)-target) < math.Abs(float64(res)-target) {
			res = value
		}
		if target < float64(root.Val) {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return res
}
