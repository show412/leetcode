# Given a binary tree, return the level order traversal of its nodes' values. (ie, from left to right, level by level).

# For example:
# Given binary tree [3,9,20,null,null,15,7],
#     3
#    / \
#   9  20
#     /  \
#    15   7
# return its level order traversal as:
# [
#   [3],
#   [9,20],
#   [15,7]
# ]
# https://leetcode.com/problems/binary-tree-level-order-traversal/description/

def level_order(root)
  return [] if root.nil?
  result = []
  q = [root]
  while !q.empty?
    # we only want values in the results
    result << q.map(&:val)
    next_lv = []
    q.each do |node|
      next_lv << node.left if node.left
      next_lv << node.right if node.right
    end
    q = next_lv
  end
  result 
end