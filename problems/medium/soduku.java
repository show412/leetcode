// https://leetcode.com/problems/valid-sudoku/submissions/
// it use the hash to store these data and every sub box is standed by the diveided by 3
class Solution {
    public boolean isValidSudoku(char[][] board) {
      Set seen = new HashSet();
      for (int i=0; i<9; ++i) {
          for (int j=0; j<9; ++j) {
              char number = board[i][j];
              if (number != '.')
                  if (!seen.add(number + " in row " + i) ||
                      !seen.add(number + " in column " + j) ||
                      !seen.add(number + " in block " + i/3 + "-" + j/3))
                      return false;
          }
      }
      return true;
  }
}