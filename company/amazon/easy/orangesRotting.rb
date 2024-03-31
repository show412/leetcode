# // https://leetcode.com/problems/rotting-oranges/
# /*
# In a given grid, each cell can have one of three values:

# the value 0 representing an empty cell;
# the value 1 representing a fresh orange;
# the value 2 representing a rotten orange.
# Every minute, any fresh orange that is adjacent (4-directionally) to a rotten orange
# becomes rotten.

# Return the minimum number of minutes that must elapse
# until no cell has a fresh orange.  If this is impossible, return -1 instead.



# Example 1:



# Input: [[2,1,1],[1,1,0],[0,1,1]]
# Output: 4
# Example 2:

# Input: [[2,1,1],[0,1,1],[1,0,1]]
# Output: -1
# Explanation:  The orange in the bottom left corner (row 2, column 0) is never rotten, because rotting only happens 4-directionally.
# Example 3:

# Input: [[0,2]]
# Output: 0
# Explanation:  Since there are already no fresh oranges at minute 0, the answer is just 0.


# Note:

# 1 <= grid.length <= 10
# 1 <= grid[0].length <= 10
# grid[i][j] is only 0, 1, or 2.
# */
# @param {Integer[][]} grid
# @return {Integer}
def oranges_rotting(grid)
  if grid.length ==1 && grid[0].length == 1
    if grid[0][0] == 1
      return -1
    end
    return 0
  end
  res = 0
  queue = []
  (0...grid.length).each do |i|
    (0...grid[0].length).each do |j|
      if grid[i][j] == 2
        point = [i,j]
        queue.push(point)
      end
    end
  end

  path = [0,1,0,-1,0]
  while queue.length != 0
    size = queue.length
    res += 1
    (0...size).each do |i|
      node = queue.shift
      r = node[0]
      c = node[1]
      (0...4).each do |k|
        r1 = r + path[k]
        c1 = c + path[k+1]
        if r1 >=0 && c1 >= 0 && r1 < grid.length && c1 < grid[0].length && grid[r1][c1] == 1
          grid[r1][c1] = 2
          queue.push([r1,c1])
        end
      end
    end
  end

  (0...grid.length).each do |i|
    (0...grid[0].length).each do |j|
      return - 1 if grid[i][j] == 1
    end
  end

  res-1

end
