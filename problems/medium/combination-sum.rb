# @param {Integer[]} candidates
# @param {Integer} target
# @return {Integer[][]}
def combination_sum(candidates, target)
  $result = []
  return $result if(candidates == nil || candidates.length == 0)
  candidates.sort!
  dfs(candidates,0,[],target)
  return $result
end

def dfs(candidates, startIndex,combinations, remain)
  if remain == 0
    $result.push(combinations.clone)
    return
  end

  i = startIndex
  while(i<candidates.length)
    if remain < candidates[i]
      break
    end
    combinations.push(candidates[i])
    dfs(candidates,i,combinations,remain - candidates[i])
    combinations.pop
    i += 1
  end
end

p combination_sum([8,7,4,3], 11)
