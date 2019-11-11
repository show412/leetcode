# https://leetcode.com/problems/critical-connections-in-a-network/
# There are n servers numbered from 0 to n-1 connected
# by undirected server-to-server connections
# forming a network where connections[i] = [a, b]
# represents a connection between servers a and b.
# Any server can reach any other server directly or indirectly through the network.

# A critical connection is a connection that,
# if removed, will make some server unable to reach some other server.

# Return all critical connections in the network in any order.



# Example 1:



# Input: n = 4, connections = [[0,1],[1,2],[2,0],[1,3]]
# Output: [[1,3]]
# Explanation: [[3,1]] is also accepted.


# Constraints:

# 1 <= n <= 10^5
# n-1 <= connections.length <= 10^5
# connections[i][0] != connections[i][1]
# There are no repeated connections.

# @param {Integer} n
# @param {Integer[][]} connections
# @return {Integer[][]}
# refer to the solution 思路参考 https://www.cnblogs.com/en-heng/p/4002658.html
# 具体过程可以参考 https://www.cnblogs.com/nullzx/p/7968110.html
# dfn[u]记录节点u在DFS过程中被遍历到的次序号，
# low[u]记录节点u或u的子树通过非父子边追溯到最早的祖先节点
# 即DFS次序号最小），那么low[u]的计算过程如下：

# 𝑙𝑜𝑤[𝑢]={min{𝑙𝑜𝑤[𝑢], 𝑙𝑜𝑤[𝑣]} (𝑢,𝑣)为树边
# 𝑙𝑜𝑤[𝑢]= min{𝑙𝑜𝑤[𝑢], 𝑑𝑓𝑛[𝑣]} (𝑢,𝑣)为回边且𝑣不为𝑢的父亲节点
# 当(u,v)为树边且low[v] >= dfn[u]时，节点u才为割点。
# 该式子的含义：以节点v为根的子树所能追溯到最早的祖先节点要么为v要么为u。
# TC is O(V+E)
def critical_connections(n, connections)
  $g = []
  (0...n).each {|i| g[i] = []}
  connections.each do |connection|
    u = connection[0]
    v = connection[1]
    $g[u].push(v)
    $g[v].push(u)
  end
  $time = 0
  $res = []
  $low = []
  $dfn = []
  # 都先赋值为最大值
  (0...n).each {|i| $dfn[i] = Float::INFINITY}
  # 从跟节点开始
  dfs(0, -1)
  res
end

def dfs(u, pre)
  $time += 1
  $dfn[u] = $time
  $low[u] = $time
  $g[u].each do |v|
    next if v == pre
    # v没有被遍历到过 并且为树边
    if $dfn[v] == Float::INFINITY
      # 因为是无向图 所以还需要反向dfs一次
      dfs(v, u)
      $low[u] = [$low[u], $low[v]].min
      # 割边是 $res.push([u,v]) if $low[v] > $dfn[u]
      # 割点 是low[v] >= dfn[u] $res.push(u)
      $res.push([u,v]) if $low[v] > $dfn[u]
    else
      # 为回边的情况
      $low[u] = [$low[u], $dfn[v]].min
    end
  end
end
