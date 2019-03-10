# @param {Integer[]} nums
# @return {Integer[][]}
def permute(nums)
  # it's used to store the result
  $result = []
  return $result if(nums == nil || nums.length == 0)
  # it's used to store the possible permute in this processing
  permute1 = []
  # it's used to store the which nums is visited
  visited = {}
  dfs(nums,permute1,visited)
  return $result
end

def dfs(nums, permute1, visited)
  if(nums.length == permute1.length)
    p $result
    $result << permute1
    return
  end
  i = 0
  # p visited
  while(i<nums.length)
    if(visited[nums[i]] == true)
      i += 1
      next
    end
    permute1 << nums[i]
    visited[nums[i]] = true
    dfs(nums, permute1, visited)
    permute1.pop
    visited[nums[i]] = false
    i += 1
  end
end
