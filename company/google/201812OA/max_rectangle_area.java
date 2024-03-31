// https://www.lintcode.com/problem/plane-maximum-rectangle/description
public class Solution {
    /**
     * @param a: the points
     * @return: return the maximum rectangle area
     */
    public int getMaximum(int[][] a) {
        // write your code here
         boolean  [][] existence = new boolean [1005][1005];
        int res = 0, nxa = 0, nya = 0, nxb = 0, nyb = 0;     // 用于计算两个点的坐标
        for(int i = 0; i < a.length; i++) {
            existence[a[i][0]] [a[i][1]] = true;    	// 标记存在的点
        }
        for(int i = 0; i < a.length; i++)
            for(int j = i+1; j < a.length; j++) {
                int x,y;
                x = a[i][0] - a[j][0];				//计算对角线两点横坐标差
                y = a[i][1] - a[j][1];				//计算对角线两点纵坐标差
                if(x * y > 0) {						//此时代表题解中的第一种情况
                    nxa = Math.max(a[i][0], a[j][0]);
                    nya = Math.min(a[i][1], a[j][1]);
                    nxb = Math.min(a[i][0], a[j][0]);
                    nyb = Math.max(a[i][1], a[j][1]);
                }
                else if(x * y < 0) {				//此时代表题解中的第二种情况
                    nxa = Math.max(a[i][0], a[j][0]);
                    nya = Math.max(a[i][1], a[j][1]);
                    nxb = Math.min(a[i][0], a[j][0]);
                    nyb = Math.min(a[i][1], a[j][1]);
                }
                if(existence[nxa][nya] == true && existence[nxb][nyb] == true) {	//如果两点存在，计算面积
                    int ans = Math.abs(x * y);
                    res = Math.max(ans, res);
                }
        }
        return res;
    }
}
