# Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

# Note:

# The solution set must not contain duplicate triplets.

# Example:

# Given array nums = [-1, 0, 1, 2, -1, -4],

# A solution set is:
# [
#   [-1, 0, 1],
#   [-1, -1, 2]
# ]
# https://leetcode.com/problems/3sum/description/

def three_sum(nums)
  sums = []
  nums.sort!
  
  for i in 0..(nums.length - 3)
      next if i > 0 && nums[i - 1] == nums[i]
      j = i + 1
      k = nums.length - 1
      while (j < k) do
          sum = nums[i] + nums[j] + nums[k]
          if sum > 0
              k -= 1
          elsif sum < 0
              j += 1
          else
              sums << [nums[i], nums[j], nums[k]]
              j += 1
              k -= 1
              j += 1 while nums[j] == nums[j - 1]
              k -= 1 while nums[k] == nums[k + 1]
          end
      end
  end
  
  return sums
end