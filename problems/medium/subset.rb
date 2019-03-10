# Given a set of distinct integers, nums, return all possible subsets (the power set).

# Note: The solution set must not contain duplicate subsets.

# Example:

# Input: nums = [1,2,3]
# Output:
# [
#   [3],
#   [1],
#   [2],
#   [1,2,3],
#   [1,3],
#   [2,3],
#   [1,2],
#   []
# ]

# @param {Integer[]} nums
# @return {Integer[][]}

def subsets(nums)
  path = []
  @result = []
  nums.sort!
  subsetsHelper(path, nums, 0, @result)
  @result
end

def subsetsHelper(path, nums, pos, result)
  result << path
  i = pos
  while(i < nums.length)
    path << nums[i]
    subsetsHelper(path, nums, i + 1, result)
    path.pop
    i += 1
  end
end

# the good solution from ruby
def subsets(nums)
  subsets = [[]]
  nums.each do |n|
    subsets += subsets.map { |s| s + [n] }
  end
  subsets
end

# very easy solution from ruby
def subsets(nums)
  (0..nums.size).flat_map{ |k| nums.combination(k).to_a }
end
