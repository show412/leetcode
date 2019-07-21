import "math"

// https://leetcode.com/problems/minimum-area-rectangle-ii/
// solution https://leetcode.com/problems/minimum-area-rectangle-ii/solution/
func minAreaFreeRect(points [][]int) float64 {
	m := make(map[[2]int]bool)
	res := math.MaxFloat64
	for i := 0; i < len(points); i++ {
		m[[2]int{points[i][0], points[i][1]}] = true
	}
	for i := 0; i < len(points); i++ {
		A := points[i]
		for j := i + 1; j < len(points); j++ {
			B := points[j]
			ABx, ABy := points[i][0]-points[j][0], points[i][1]-points[j][1]
			for k := j + 1; k < len(points); k++ {
				C := points[k]
				ACx, ACy := points[i][0]-points[k][0], points[i][1]-points[k][1]
				if ABx*ACx+ABy*ACy != 0 {
					continue
				}
				D := [2]int{B[0] - A[0] + C[0], B[1] - A[1] + C[1]}
				if _, ok := m[D]; !ok {
					continue
				}
				area := length(ABx, ABy) * length(ACx, ACy)
				res = math.Min(res, area)
			}
		}
	}
	if res == math.MaxFloat64 {
		return 0
	}
	return res
}

func length(x, y int) float64 {
	return math.Sqrt(float64(x*x + y*y))
}

/*
import java.awt.Point;

class Solution {
    public double minAreaFreeRect(int[][] points) {
        int N = points.length;
        Point[] A = new Point[N];
        for (int i = 0; i < N; ++i)
            A[i] = new Point(points[i][0], points[i][1]);

        Map<Integer, Map<Point, List<Point>>> seen = new HashMap();
        for (int i = 0; i < N; ++i)
            for (int j = i+1; j < N; ++j) {
                // center is twice actual to keep integer precision
                Point center = new Point(A[i].x + A[j].x, A[i].y + A[j].y);

                int r2 = (A[i].x - A[j].x) * (A[i].x - A[j].x);
                r2 += (A[i].y - A[j].y) * (A[i].y - A[j].y);
                if (!seen.containsKey(r2))
                    seen.put(r2, new HashMap<Point, List<Point>>());
                if (!seen.get(r2).containsKey(center))
                    seen.get(r2).put(center, new ArrayList<Point>());
                seen.get(r2).get(center).add(A[i]);
            }

        double ans = Double.MAX_VALUE;
        for (Map<Point, List<Point>> info: seen.values()) {
            for (Point center: info.keySet()) {  // center is twice actual
                List<Point> candidates = info.get(center);
                int clen = candidates.size();
                for (int i = 0; i < clen; ++i)
                    for (int j = i+1; j < clen; ++j) {
                        Point P = candidates.get(i);
                        Point Q = candidates.get(j);
                        Point Q2 = new Point(center);
                        Q2.translate(-Q.x, -Q.y);
                        double area = P.distance(Q) * P.distance(Q2);
                        if (area < ans)
                            ans = area;
                    }
            }
        }

        return ans < Double.MAX_VALUE ? ans : 0;
    }
}
*/
