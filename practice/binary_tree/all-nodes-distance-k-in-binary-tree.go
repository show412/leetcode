// https://leetcode.com/problems/all-nodes-distance-k-in-binary-tree/
// We are given a binary tree (with root node root), a target node, and an integer value K.

// Return a list of the values of all nodes that have a distance K from the target node.
// The answer can be returned in any order.

// Example 1:

// Input: root = [3,5,1,6,2,0,8,null,null,7,4], target = 5, K = 2

// Output: [7,4,1]

// Explanation:
// The nodes that are a distance 2 from the target node (with value 5)
// have values 7, 4, and 1.

// Note that the inputs "root" and "target" are actually TreeNodes.
// The descriptions of the inputs above are just serializations of these objects.

// Note:

// The given tree is non-empty.
// Each node in the tree has unique values 0 <= node.val <= 500.
// The target node is a node in the tree.
// 0 <= K <= 1000.
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func distanceK(root *TreeNode, target *TreeNode, K int) []int {
	if K == 0 {
		return []int{target.Val}
	}
	parent := make(map[*TreeNode]*TreeNode, 0)
	visited := make(map[*TreeNode]bool, 0)
	getParent(root, nil, parent)
	stack := make([]*TreeNode, 0)
	// 先放target 后放nil level初始化就应该是1 因为达到K的时候 当前的符合的level还应该在stack里
	// 用nil进行分层
	stack = append(stack, target)
	stack = append(stack, nil)
	level := 1
	visited[target] = true
	visited[nil] = true
	res := make([]int, 0)
	for len(stack) != 0 {
		node := stack[0]
		stack = stack[1:]
		if node == nil {
			if level == K {
				for i := 0; i < len(stack); i++ {
					if stack[i] != nil {
						res = append(res, stack[i].Val)
					}
				}
				return res
			}
			stack = append(stack, nil)
			level++
		} else {
			if _, ok := visited[node.Left]; !ok && node.Left != nil {
				visited[node.Left] = true
				stack = append(stack, node.Left)
			}
			if _, ok := visited[node.Right]; !ok && node.Right != nil {
				visited[node.Right] = true
				stack = append(stack, node.Right)
			}
			pNode := parent[node]
			if _, ok := visited[pNode]; !ok {
				visited[pNode] = true
				stack = append(stack, pNode)
			}
			// 不能写在这 why? 还没搞明白
			// stack = append(stack, nil)
		}
	}
	return res
}

func getParent(node *TreeNode, pNode *TreeNode, parent map[*TreeNode]*TreeNode) {
	parent[node] = pNode
	if node.Left != nil {
		getParent(node.Left, node, parent)
	}
	if node.Right != nil {
		getParent(node.Right, node, parent)
	}
	return
}
