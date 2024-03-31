# // https://leetcode.com/problems/4sum/
# /*
# Given an array nums of n integers and an integer target,
# are there elements a, b, c, and d in nums such that a + b + c + d = target?
# Find all unique quadruplets in the array which gives the sum of target.

# Note:

# The solution set must not contain duplicate quadruplets.

# Example:

# Given array nums = [1, 0, -1, 0, -2, 2], and target = 0.

# A solution set is:
# [
#   [-1,  0, 0, 1],
#   [-2, -1, 1, 2],
#   [-2,  0, 0, 2]
# ]
# */
# @param {Integer[]} nums
# @param {Integer} target
# @return {Integer[][]}
def four_sum(nums, target)
  res = []
  n = nums.length
  nums.sort!
  (0..n-3).each do |i|
    next if i > 0 && nums[i] == nums[i-1]
    (i+1..n-2).each do |j|
      next if j > i+1 && nums[j] == nums[j-1]
      k = j+1
      l = n-1
      while k < l
        sum = nums[i] + nums[j] + nums[k] + nums[l]
        if sum == target 
          res.push([nums[i], nums[j], nums[k], nums[l]])
          k += 1
          l -= 1
          k += 1 while nums[k] == nums[k-1]
          l -= 1 while nums[l] == nums[l+1]
        elsif sum < target
          k += 1
        else 
          l -= 1
        end
      end

    end
  end   
  res 
end