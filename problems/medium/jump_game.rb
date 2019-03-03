# @param {Integer[]} nums
# @return {Boolean}
def can_jump(nums)
  # state, it stands for could it jump to this postion
  f = []
  i = 0
  while(i<nums.length)
    f[i] = false
    i += 1
  end
  f[0] = true
  i = 1
  while(i<nums.length)
    j = 0
    while(j < i)
      if(f[j] == true && j + nums[j] >= i)
        f[i] = true
        break
      end
      j += 1
    end
    i += 1
  end
  f[nums.length - 1]
end
