# @param {Integer[]} nums
# @return {Integer[]}
def majority_element(nums)
  #     quick sort first then travers the array
  # quick sort Time complexity is O(nlogn) space complexity is O(1)
  return [] if nums.nil? || nums.empty?
  nums.sort!
  length = nums.length
  size = length / 3
  result = []
  i = 0
  j = 0
  count = 1
  while(j < length)
    j += 1
    if nums[i] == nums[j]
      count += 1
    else
      result.push(nums[i]) if count > size
      i = j
      count = 1
    end
  end
  result
end

# time complexity is O(n) space complexity is O(1)
def majority_element(nums)
  return [] if nums.nil? || nums.empty?

  length = nums.length
  i = 0
  candidate1 = nil
  candidate2 = nil
  count1 = 0
  count2 = 0
  while(i<length)
    if(nums[i] == candidate1)
      count1 += 1
    elsif(nums[i] == candidate2)
      count2 += 1
    elsif(count1 == 0)
      candidate1 = nums[i]
    elsif(count2 == 0)
      candidate2 = nums[i]
    else
      count1 -= 1
      count2 -= 1
    end
    i += 1
  end
  count1 = 0
  count2 = 0
  i = 0
  while(i<length)
    count1 += 1 if(nums[i] == candidate1)
    count2 += 1 if(nums[i] == candidate2)
    i += 1
  end
  count1 > count2 ? candidate1 : candidate2
end
