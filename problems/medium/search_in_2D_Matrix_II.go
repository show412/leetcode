// https://leetcode.com/problems/search-a-2d-matrix-ii/
func searchMatrix(matrix [][]int, target int) bool {
  if len(matrix) == 0 {
    return false
  }
  if len(matrix[0]) == 0 {
    return false
  }
   r := len(matrix)-1
   c := 0
   // count := 0
  
  for r >=0 && c < len(matrix[0]) {
    // fmt.Println(matrix[r][c])
    if target == matrix[r][c] {
      return true
    }
    if target < matrix[r][c] {
      r--
    } else {
      c++
    }
  }
  return false
}