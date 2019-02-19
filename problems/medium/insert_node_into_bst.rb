# Definition for a binary tree node.
# class TreeNode
#     attr_accessor :val, :left, :right
#     def initialize(val)
#         @val = val
#         @left, @right = nil, nil
#     end
# end

# @param {TreeNode} root
# @param {Integer} val
# @return {TreeNode}
# https://leetcode.com/problems/insert-into-a-binary-search-tree/
# there is multiple ways to do, insert the node into bottom is easier
def insert_into_bst(root, val)
  return val if root.nil?
  if val.val < root.val do
    root.left = insert_into_bst(root.left, val)
  end
  if val.val > root.val do
    root.right = insert_into_bst(root.right, val)
  end
  return root
end