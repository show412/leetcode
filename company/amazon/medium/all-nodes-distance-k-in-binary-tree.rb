# // https://leetcode.com/problems/all-nodes-distance-k-in-binary-tree/
# /*
# We are given a binary tree (with root node root), a target node, and an integer value K.

# Return a list of the values of all nodes that have a distance K from the target node.
# The answer can be returned in any order.



# Example 1:

# Input: root = [3,5,1,6,2,0,8,null,null,7,4], target = 5, K = 2

# Output: [7,4,1]

# Explanation:
# The nodes that are a distance 2 from the target node (with value 5)
# have values 7, 4, and 1.



# Note that the inputs "root" and "target" are actually TreeNodes.
# The descriptions of the inputs above are just serializations of these objects.


# Note:

# The given tree is non-empty.
# Each node in the tree has unique values 0 <= node.val <= 500.
# The target node is a node in the tree.
# 0 <= K <= 1000.

# Definition for a binary tree node.
# class TreeNode
#     attr_accessor :val, :left, :right
#     def initialize(val)
#         @val = val
#         @left, @right = nil, nil
#     end
# end

# @param {TreeNode} root
# @param {TreeNode} target
# @param {Integer} k
# @return {Integer[]}

# Definition for a binary tree node.
# class TreeNode
#     attr_accessor :val, :left, :right
#     def initialize(val)
#         @val = val
#         @left, @right = nil, nil
#     end
# end

# @param {TreeNode} root
# @param {TreeNode} target
# @param {Integer} k
# @return {Integer[]}
def distance_k(root, target, k)
  return [target.val] if k == 0
  parent = {}
  visited = {}
  parent = getParent(root, nil, parent)
  stack = [target, nil]
  visited[target] = true
  visited[nil] = true
  level = 1
  res = []
  while stack.length != 0
    node = stack.shift
    if node == nil
      if level == k
        (0..stack.length).each do |i|
          res.push(stack[i].val) if stack[i] != nil
        end
        return res
      end
      stack.push(nil)
      level += 1
    else
      if visited[node.left] != true && node.left != nil
        visited[node.left] = true
        stack.push(node.left)
      end
      if visited[node.right] != true && node.right != nil
        visited[node.right] = true
        stack.push(node.right)
      end
      if visited[parent[node]] != true
        visited[parent[node]] = true
        stack.push(parent[node])
      end

    end
  end
  res
end

def getParent(node, pNode, parent)
  parent[node] = pNode
  parent = getParent(node.left, node, parent) if node.left != nil
  parent = getParent(node.right, node, parent) if node.right != nil
  parent
end
