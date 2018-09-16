# Given a matrix of m x n elements (m rows, n columns), return all elements of the matrix in spiral order.

# Example 1:

# Input:
# [
#  [ 1, 2, 3 ],
#  [ 4, 5, 6 ],
#  [ 7, 8, 9 ]
# ]
# Output: [1,2,3,6,9,8,7,4,5]
# Example 2:

# Input:
# [
#   [1, 2, 3, 4],
#   [5, 6, 7, 8],
#   [9,10,11,12]
# ]
# Output: [1,2,3,4,8,12,11,10,9,5,6,7]

# https://leetcode.com/problems/spiral-matrix/description/

# @param {Integer[][]} matrix
# @return {Integer[]}
def spiral_order(matrix)
  out_put_func(matrix).flatten.compact
end

def out_put_func(array)
  return [] if array.length == 0
  i = array.length
  j = array[0].length
  return array[0] if i == 1
  return array.flatten if j == 1
  out_put = []
  z1 = 0
  while (z1 < j) do 
    out_put << array[0][z1]
    z1 += 1
  end
  z2 = 1
  while (z2 < i-1) do 
    out_put << array[z2][j-1]
    z2 += 1
  end
  z3 = j - 1
  while (z3 > 0) do
    out_put << array[i-1][z3]
    z3 -= 1
  end
  z4 = i - 1
  while (z4 > 0) do
    out_put << array[z4][0]
    z4 -= 1
  end
  out_put << out_put_func(convert_array(array))
  out_put
end

def convert_array(array)
  c_array = []
  i = array.length
  z = 1
  while(z<i-1) do 
    array[z].pop
    array[z].shift
    c_array << array[z]
    z += 1
  end
  c_array
end