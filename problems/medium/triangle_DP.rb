# @param {Integer[][]} triangle
# @return {Integer}
def minimum_total(triangle)
  #   f is used to set the min path from botom to top
    i=0
    f = []
    # state: f stands that the min path from i to j
    while (i<triangle.length)
      f.push({})
      i += 1
    end
    # initalize: for the last level the min path
    i = 0
    while (i<triangle.length)
      f[triangle.length - 1][i] = triangle[triangle.length - 1][i]
      i += 1
    end
    # func: from bottom to top
    i = triangle.length - 2
    while (i>=0)
      j = 0
      while(j<=i)
        f[i][j] = [f[i+1][j],f[i+1][j+1]].min + triangle[i][j]
        j += 1
      end
      i -= 1
    end
    f[0][0]
  end
