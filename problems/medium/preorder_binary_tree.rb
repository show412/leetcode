# https://leetcode.com/problems/binary-tree-preorder-traversal/submissions/
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
def preorder_traversal(root)
  @result = []
  traverse(root)
  @result
end

def traverse(root)
 return if root.nil?
 @result.push(root.val)
 traverse(root.left)
 traverse(root.right)
end