// https://leetcode.com/problems/all-nodes-distance-k-in-binary-tree/
/*
We are given a binary tree (with root node root), a target node, and an integer value K.

Return a list of the values of all nodes that have a distance K from the target node.
The answer can be returned in any order.



Example 1:

Input: root = [3,5,1,6,2,0,8,null,null,7,4], target = 5, K = 2

Output: [7,4,1]

Explanation:
The nodes that are a distance 2 from the target node (with value 5)
have values 7, 4, and 1.



Note that the inputs "root" and "target" are actually TreeNodes.
The descriptions of the inputs above are just serializations of these objects.


Note:

The given tree is non-empty.
Each node in the tree has unique values 0 <= node.val <= 500.
The target node is a node in the tree.
0 <= K <= 1000.
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//  先dfs记node 的父子关系, 再bfs 通过加 null 节点来分层 这是 tree 的一个 common 的做法
func distanceK(root *TreeNode, target *TreeNode, K int) []int {
	parent := make(map[*TreeNode]*TreeNode, 0)
	seen := make(map[*TreeNode]bool, 0)
	dfs(root, nil, parent)
	queue := make([]*TreeNode, 0)
	res := make([]int, 0)
	// 必须得先放 nil 这样先加1
	// 如果把 nil 放后面 dist 初始化是1 当K == 0的时候 就死循环了
	queue = append(queue, nil)
	queue = append(queue, target)
	dist := 0
	seen[target] = true
	seen[nil] = true
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			if dist == K {
				for _, n := range queue {
					if n != nil {
						res = append(res, n.Val)
					}
				}
				return res
			}
			queue = append(queue, nil)
			dist++
		} else {
			if _, ok := seen[node.Left]; !ok {
				seen[node.Left] = true
				queue = append(queue, node.Left)
			}
			if _, ok := seen[node.Right]; !ok {
				seen[node.Right] = true
				queue = append(queue, node.Right)
			}
			par := parent[node]
			if _, ok := seen[par]; !ok {
				seen[par] = true
				queue = append(queue, par)
			}
		}
	}
	return res
}

func dfs(node *TreeNode, par *TreeNode, parent map[*TreeNode]*TreeNode) {
	if node != nil {
		parent[node] = par
		dfs(node.Left, node, parent)
		dfs(node.Right, node, parent)
	}
}
