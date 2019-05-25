# Given an integer array with no duplicates. A maximum tree building on this array is defined as follow:

# The root is the maximum number in the array.
# The left subtree is the maximum tree constructed from left part subarray divided by the maximum number.
# The right subtree is the maximum tree constructed from right part subarray divided by the maximum number.
# Construct the maximum tree by the given array and output the root node of this tree.

# Example 1:
# Input: [3,2,1,6,0,5]
# Output: return the tree root node representing the following tree:

#       6
#     /   \
#    3     5
#     \    /
#      2  0
#        \
#         1
# Note:
# The size of the given array will be in the range [1,1000].
# https://leetcode.com/problems/maximum-binary-tree/description/

# Definition for a binary tree node.
class TreeNode
    attr_accessor :val, :left, :right
    def initialize(val)
        @val = val
        @left, @right = nil, nil
    end
end

# @param {Integer[]} nums
# @return {TreeNode}
def construct_maximum_binary_tree(nums)
  tree_node = TreeNode.new(nums.max)
  if nums.length == 1
    return tree_node
  end
  max_at = nums.index(nums.max)
  left_nums = nums.slice(0, max_at)
  right_nums = nums.slice(max_at+1, nums.length)
  tree_node.left = construct_maximum_binary_tree(left_nums) if left_nums != []
  tree_node.right = construct_maximum_binary_tree(right_nums) if right_nums != []
  tree_node
end

# Stack 的方式
# java
public class Solution {
    public TreeNode maxTree(int[] A) {
        if (A == null || A.length == 0) return null;
        Stack<TreeNode> stack = new Stack<>();
        for (int i = 0; i < A.length; i++) {
            # //遍历A的每个元素，创造结点node
            TreeNode node = new TreeNode(A[i]);
            # //将stack中小于当前结点的结点都pop出来，存为当前结点的左子树
            while (!stack.isEmpty() && node.val >= stack.peek().val) node.left = stack.pop();
            # //如果stack仍非空，剩下的结点一定大于当前结点，
            # 所以将当前结点存为stack中结点的右子树；
            # 而stack中结点本来的右子树之前已经存为当前结点的左子树了
            if (!stack.isEmpty()) stack.peek().right = node;
            # //stack中存放结点的顺序为：底部为完整的max tree，
            # 从下向上是下一层孩子结点的备份，顶部是当前结点的备份
            stack.push(node);
        }
        TreeNode root = stack.pop();
        while (!stack.isEmpty()) root = stack.pop();
        return root;
    }
}