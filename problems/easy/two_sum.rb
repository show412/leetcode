# Given an array of integers, return indices of the two numbers such that they add up to a specific target.

# You may assume that each input would have exactly one solution, and you may not use the same element twice.

# Example:

# Given nums = [2, 7, 11, 15], target = 9,

# Because nums[0] + nums[1] = 2 + 7 = 9,
# return [0, 1].

# @param {Integer[]} nums
# @param {Integer} target
# @return {Integer[]}
# my solution
def two_sum(nums, target)
  index1 = nil
  index2 = nil
  nums.map.with_index do |value, index|
    index1 = index
    target_index = nums.rindex(target - value)
    if !target_index.nil? && target_index > index
        index2 = target_index
        break
    else
        index1 = nil
        index2 = nil
    end
  end
  [index1, index2]
end

# best solution
def two_sum(nums, target)
  memo = {}
  nums.each_with_index do |num, i|
    remainder = target - num
    return [memo[remainder], i] if memo[remainder]
    memo[num] = i
  end
  raise ArgumentError, "Expected one valid solution"
end
