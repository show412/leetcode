# @param {Integer} n
# @return {Integer}
def climb_stairs(n)
  return 1 if n == 0 
  # f[n] stands for how many possiblities to climb to n-th stairs
  f = {}
  f[-1] = 0
  f[0] = 1
  f[1] = 2
  i = 1
  while (i<=n)
    f[i] = f[i-1] + f[i-2]
    i += 1
  end
  f[n]
end