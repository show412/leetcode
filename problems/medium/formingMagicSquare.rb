# Complete the formingMagicSquare function below.
def formingMagicSquare(s)
 allPossible = [
    [8, 1, 6, 3, 5, 7, 4, 9, 2],
		[6, 1, 8, 7, 5, 3, 2, 9, 4],
		[4, 9, 2, 3, 5, 7, 8, 1, 6],
		[2, 9, 4, 7, 5, 3, 6, 1, 8],
		[8, 3, 4, 1, 5, 9, 6, 7, 2],
		[4, 3, 8, 9, 5, 1, 2, 7, 6],
		[6, 7, 2, 1, 5, 9, 8, 3, 4],
		[2, 7, 6, 9, 5, 1, 4, 3, 8]
 ]
 flatS = [s[0][0], s[0][1], s[0][2], s[1][0], s[1][1], s[1][2], s[2][0], s[2][1], s[2][2]]
 adjustNum = []
 i = 0
 sum = 0
 while(i<8)
  sum = 0
  j = 0
  while(j<9)
    sum += (allPossible[i][j] - flatS[j]).abs
    j +=1
  end
  adjustNum.push(sum)
  i +=1
 end
 adjustNum.min
end
