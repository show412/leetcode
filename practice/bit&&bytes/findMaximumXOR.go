// https://leetcode.com/problems/maximum-xor-of-two-numbers-in-an-array/
/*
Given a non-empty array of numbers, a0, a1, a2, … , an-1, where 0 ≤ ai < 231.

Find the maximum result of ai XOR aj, where 0 ≤ i, j < n.

Could you do this in O(n) runtime?

Example:

Input: [3, 10, 5, 25, 2, 8]

Output: 28

Explanation: The maximum result is 5 ^ 25 = 28.

The important thing to note is, we don't need to return a pair of indices of a result but just need a result of XOR.

Because a number of bits in integer is constant (in Go, leetcode has 32-bit integer environment), we can first build a Trie which has all possible bits (0 or 1) in each digit (from 31 to 0 th) among all elements in an array.

Then we can again iterate through the array and for each element, we will traverse the Trie so that XOR of each digit should have maximum value. Obviously, the more significant bit has 1, a result of XOR becomes bigger.

refer to https://leetcode.com/problems/maximum-xor-of-two-numbers-in-an-array/discuss/130427/()-92

*/
func findMaximumXOR(nums []int) int {
	root := NewNode()

	// build trie
	// 为所有数构建在一个trie数上 有些0，1节点是复用的
	// 那么这颗树包含了数组中所有的数字，
	// 从根节点到叶子节点的一条路径表示数组中的一个十进制数的二进制编码。
	for _, num := range nums {
		cur := root
		for i := 31; i >= 0; i-- {
			flag := (num >> uint(i)) & 1
			if cur.Children[flag] == nil {
				cur.Children[flag] = NewNode()
			}
			cur = cur.Children[flag]
		}
	}

	// search maximum
	max := 0
	for _, num := range nums {
		// 从root开始遍历
		cur, val := root, 0
		for i := 31; i >= 0; i-- {
			flag := (num >> uint(i)) & 1
			switch {
			// 当前位为1 则与0异或结果最大； 当前位为0 则与1异或结果最大
			case flag == 1 && cur.Children[0] != nil, flag == 0 && cur.Children[1] != nil:
				// 把1左移挪到最大的位上
				val += 1 << uint(i)
				cur = cur.Children[1 & ^flag] // just reversing flag; ^flag means flag的非
			default:
				// 如果最大的分支不存在只能选择存在的分支
				cur = cur.Children[flag]
			}
		}
		if val > max {
			max = val
		}
	}
	return max
}

type TNode struct {
	Children [2]*TNode
}

func NewNode() *TNode {
	return &TNode{Children: [2]*TNode{nil, nil}}
}

// https://github.com/grandyang/leetcode/issues/421
// another find the max xor result solution
