// https://leetcode.com/problems/search-a-2d-matrix/
// sss
func searchMatrix(matrix [][]int, target int) bool {
  if len(matrix) == 0 ||len(matrix[0]) == 0 || target < matrix[0][0] {
    return false
  }
  line_count := len(matrix)
  row_count := len(matrix[0])
//   search which row, find the last position which is smaller than target in row[0]
  var start = 0
  var end = line_count - 1
  var mid = 0
  var row_position = 0
  for {
    if start + 1 >= end {
      break
    }
    mid = start + (end - start)/2
    if matrix[mid][0] == target {
      return true
    }else if matrix[mid][0] < target {
      start = mid
    }else if matrix[mid][0] > target {
      end = mid
    }    
  }
	
	// find the last position where target is bigger that the row[0]
	// so it should put the end at the last to rewrite the row_position
  if matrix[start][0] < target {
    row_position = start
  }
  
  if matrix[end][0] < target {
    row_position = end
  }
  
  if matrix[start][0] == target {
    return true
  }
  if matrix[end][0] == target {
    return true
  }
  // fmt.Println(start)
  // fmt.Println(end)
	// fmt.Println(row_position)
	
//   find which line the target
  start = 0
  end = row_count - 1
  mid = 0
  for {
    if start + 1 >= end {
      break
    }
    mid = start + (end - start)/2
    if matrix[row_position][mid] == target {
      return true
    }else if matrix[row_position][mid] < target {
      start = mid
    }else if matrix[row_position][mid] > target {
      end = mid
    }    
  }
  
  if matrix[row_position][end] == target {
    return true
  }
  if matrix[row_position][start] == target {
    return true
  }
  return false
  
}