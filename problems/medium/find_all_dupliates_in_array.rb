# Given an array of integers, 1 ≤ a[i] ≤ n (n = size of array), some elements appear twice and others appear once.

# Find all the elements that appear twice in this array.

# Could you do it without extra space and in O(n) runtime?

# Example:
# Input:
# [4,3,2,7,8,2,3,1]

# Output:
# [2,3]
# https://leetcode.com/problems/find-all-duplicates-in-an-array/description/

# @param {Integer[]} nums
# @return {Integer[]}
def find_duplicates(nums)
  res = []
  return res if nums.nil? || nums.length == 0
  length = nums.length
  i = 0
  while (i<length) do
    # p "i is #{i}"
    num = nums[i]
    if (nums[i] == i + 1)
      i += 1
      next
    end
    if (num == -1)
      i +=1
      next
    end
    if (nums[num-1] == num )
      nums[i] = -1
      res.push(num)
      # p "res is #{res}"
      i += 1
      next
    end
    temp = nums[i]
    nums[i] = nums[num-1]
    nums[num-1] = temp
    # p "nums is #{nums}"
  end
  return res
end
