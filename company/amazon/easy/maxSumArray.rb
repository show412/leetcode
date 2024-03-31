# // https://leetcode.com/problems/maximum-subarray/
# /*
# Given an integer array nums, find the contiguous subarray
# (containing at least one number)
# which has the largest sum and return its sum.

# Example:

# Input: [-2,1,-3,4,-1,2,1,-5,4],
# Output: 6
# Explanation: [4,-1,2,1] has the largest sum = 6.
# Follow up:

# If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach,
# which is more subtle.
# @param {Integer[]} nums
# @return {Integer}
def max_sub_array(nums)
  return 0 if nums.length == 0
  return nums[0] if nums.length == 1
  max = nums[0]
  maxSum = nums[0]
  (1...nums.length).each do |i|
    if max >= 0
      max += nums[i]
    else
      max = nums[i]
    end
    maxSum = max if maxSum < max
  end
  return maxSum
end
