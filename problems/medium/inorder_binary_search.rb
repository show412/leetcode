# https://leetcode.com/problems/binary-tree-inorder-traversal/submissions/
# Definition for a binary tree node.
# class TreeNode
#     attr_accessor :val, :left, :right
#     def initialize(val)
#         @val = val
#         @left, @right = nil, nil
#     end
# end

# @param {TreeNode} root
# @return {Integer[]}
# it's Divide & Conquer to resolve the problem
def inorder_traversal(root)
  result = []
  return result if root.nil?
  left = inorder_traversal(root.left)
  right = inorder_traversal(root.right)
  result = left + [root.val] + right
  result
end