# Definition for a binary tree node.
# class TreeNode
#     attr_accessor :val, :left, :right
#     def initialize(val)
#         @val = val
#         @left, @right = nil, nil
#     end
# end

# @param {TreeNode} root
# @param {TreeNode} p
# @param {TreeNode} q
# @return {TreeNode}
# https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/description/

def lowest_common_ancestor(root, p, q)
  return nil if root.nil?
  return p if in_node?(p, q)
  return q if in_node?(q, p)
  until root.nil?
      return root if [nil, p, q].index(root)
      if in_node?(root.left, p) && in_node?(root.left, q)
          root = root.left
          next
      end
      if in_node?(root.right, p) && in_node?(root.right, q)
          root = root.right
          next
      end
      return root
  end
end

def in_node?(parent, node)
  return false if parent.nil?
  return true if parent == node
  in_node?(parent.left, node) || in_node?(parent.right, node)
end