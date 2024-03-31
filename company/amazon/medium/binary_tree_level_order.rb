# // https://leetcode.com/problems/binary-tree-level-order-traversal/
# /*
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
# */
# /**
#  * Definition for a binary tree node.
#  * type TreeNode struct {
#  *     Val int
#  *     Left *TreeNode
#  *     Right *TreeNode
#  * }
#  */

# Definition for a binary tree node.
# class TreeNode
#     attr_accessor :val, :left, :right
#     def initialize(val)
#         @val = val
#         @left, @right = nil, nil
#     end
# end

# @param {TreeNode} root
# @return {Integer[][]}
def level_order(root)
  return [] if root == nil
  res = []
  level = [root]
  until level.empty?
    next_level = []
    cur_level = []
    (0...level.length).each do |i|
      cur_level.push(level[i].val)
      next_level.push(level[i].left) if level[i].left != nil
      next_level.push(level[i].right) if level[i].right != nil
    end
    res.push(cur_level)
    level = next_level
  end
  res
end