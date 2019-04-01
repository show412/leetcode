// 假设buff[i][j]的含义为：将第i个数改成j的调整代价。如果改成j之后不符合调整的目标则buff[i][j]=INT_MAX。

// 初始化缓存将第0个数改为j的代价是

// buff[0][j] = abs(A[0]-j);
// 1
// 为了符合调整目标，将第i个数改成j的话，第i-1个数k的可取值范围是[j-target, j+target]。在符合调整目标的情况下的状态转移方程为：

// buff[i][j] = min(buff[i][j], buff[i-1][k]+abs(j-A[i]));

class Solution {
public:
    /**
     * @param A: An integer array.
     * @param target: An integer.
     */
    int MinAdjustmentCost(vector<int> A, int target) {
        // write your code here
        const int len = A.size();
        int buff[len][101];
        for (int j = 1; j <= 100; j++)
        {
            buff[0][j] = abs(j-A[0]);
        }
        for (int i = 1; i < len; i++)
        {
            for (int j = 1; j <= 100; j++)
            {
                buff[i][j] = INT_MAX;
                int lo = max(1, j - target);
                int hi = min(100, j + target);
                for (int k = lo; k <= hi; k++)
                {
                    buff[i][j] = min(buff[i][j], buff[i-1][k]+abs(j-A[i]));
                }
            }
        }
        int res = INT_MAX;
        for (int i = 1; i <= 100; i++)
        {
            res = min(res, buff[len-1][i]);
        }
        return res;
    }
};
