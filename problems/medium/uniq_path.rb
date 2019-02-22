# https://leetcode.com/problems/unique-paths/
# @param {Integer} m
# @param {Integer} n
# @return {Integer}
def unique_paths(m, n)
  return 1 if (m == 0 || n == 0)
  i=0
  # define the state
  # f is the possible uniq paths that arrive at this point. 
  f = []
  while (i<m)
    f.push({})
    i += 1
  end
  # initalize
  f[0][0] = 1
  i = 0
  while(i<m)
    f[i][0] = 1
    i += 1
  end
  i = 0
  while(i<n)
    f[0][i] = 1
    i += 1
  end
  # function
  i = 1
  while(i<m)
    j = 1
    while(j<n)
      f[i][j] = f[i-1][j] + f[i][j-1]
      j += 1
    end
    i += 1
  end
  # result
  f[m-1][n-1]
end