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
p four_sum([4,-9,-2,-2,-7,9,9,5,10,-10,4,5,2,-4,-2,4,-9,5],-13)
