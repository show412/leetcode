// https://leetcode.com/problems/most-stones-removed-with-same-row-or-column/
/*
On a 2D plane, we place stones at some integer coordinate points.
Each coordinate point may have at most one stone.

Now, a move consists of removing a stone
that shares a column or row with another stone on the grid.

What is the largest possible number of moves we can make?

Example 1:

Input: stones = [[0,0],[0,1],[1,0],[1,2],[2,1],[2,2]]
Output: 5
Example 2:

Input: stones = [[0,0],[0,2],[1,1],[2,0],[2,2]]
Output: 3
Example 3:

Input: stones = [[0,0]]
Output: 0

1 <= stones.length <= 1000
0 <= stones[i][j] < 10000
*/
// 这种题第一反应也是并查集 并查集的关键就是找关系

func removeStones(stones [][]int) int {

}

/*
there is a find union solution
*/
class Solution {
    public int removeStones(int[][] stones) {
        int N = stones.length;
        DSU dsu = new DSU(20000);

        for (int[] stone: stones)
            dsu.union(stone[0], stone[1] + 10000);

        Set<Integer> seen = new HashSet();
        for (int[] stone: stones)
            seen.add(dsu.find(stone[0]));

        return N - seen.size();
    }
}

class DSU {
    int[] parent;
    public DSU(int N) {
        parent = new int[N];
        for (int i = 0; i < N; ++i)
            parent[i] = i;
    }
    public int find(int x) {
        if (parent[x] != x) parent[x] = find(parent[x]);
        return parent[x];
    }
    public void union(int x, int y) {
        parent[find(x)] = find(y);
    }
}


    Map<Integer, Integer> f = new HashMap<>();
    int islands = 0;

    public int removeStones(int[][] stones) {
        for (int i = 0; i < stones.length; ++i)
            union(stones[i][0], ~stones[i][1]);
        return stones.length - islands;
    }

    public int find(int x) {
        if (f.putIfAbsent(x, x) == null)
            islands++;
        if (x != f.get(x))
            f.put(x, find(f.get(x)));
        return f.get(x);
    }

    public void union(int x, int y) {
        x = find(x);
        y = find(y);
        if (x != y) {
            f.put(x, y);
            islands--;
        }
    }


// this is a DFS to refer

func removeStones(stones [][]int) int {
	visited := make(map[int]map[int]bool)
	xm := make(map[int][]int)
	ym := make(map[int][]int)
	for _, s := range stones {
		x, y := s[0], s[1]
		if _, ok := xm[x]; !ok {
			xm[x] = make([]int, 0)
		}
		xm[x] = append(xm[x], y)
		if _, ok := ym[y]; !ok {
			ym[y] = make([]int, 0)
		}
		ym[y] = append(ym[y], x)
		if _, ok := visited[x]; !ok {
			visited[x] = make(map[int]bool)
		}
		visited[x][y] = false
	}
	count := len(stones)
	for _, s := range stones {
		if !visited[s[0]][s[1]] {
			dfs(s[0], s[1], xm, ym, visited)
			count--
		}
	}
	return count
}

func dfs(x, y int, xm, ym map[int][]int, visited map[int]map[int]bool) {
	visited[x][y] = true
	for i := 0; i < len(xm[x]); i++ {
		if !visited[x][xm[x][i]] {
			dfs(x, xm[x][i], xm, ym, visited)
		}
	}
	for j := 0; j < len(ym[y]); j++ {
		if !visited[ym[y][j]][y] {
			dfs(ym[y][j], y, xm, ym, visited)
		}
	}
}
