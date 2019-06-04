import (
	"fmt"
	"math"
	"sort"
)

// Given a set of points in the xy-plane, determine the minimum area of a rectangle formed from these points, with sides parallel to the x and y axes.

// If there isn't any rectangle, return 0.

// Example 1:

// Input: [[1,1],[1,3],[3,1],[3,3],[2,2]]
// Output: 4
// Example 2:

// Input: [[1,1],[1,3],[3,1],[3,3],[4,1],[4,3]]
// Output: 2
// https://leetcode.com/problems/minimum-area-rectangle/
// this solution is time limitted exceed
func minAreaRect(points [][]int) int {
	if len(points) < 4 {
		return 0
	}
	res := []int{}
	subPoints := [][]int{}
	dfs(&points, &subPoints, 0, &res)
	min := int(^uint(0) >> 1)
	fmt.Println(res)
	if len(res) == 0 {
		return 0
	}
	for i := 0; i < len(res); i++ {
		if min > res[i] {
			min = res[i]
		}
	}
	return min
}

func dfs(points *[][]int, subPoints *[][]int, start int, res *[]int) {
	if len(*subPoints) == 4 {
		// areaRect is for all possible point to combine a rect
		// area := areaRect(*subPoints)
		// parollAreaRect is for those x,y is paroll to x-axis and y-axis
		area := parollAreaRect(*subPoints)
		if area > 0 {
			*res = append(*res, area)
			return
		}
		return
	}

	for i := start; i < len(*points); i++ {
		*subPoints = append(*subPoints, (*points)[i])
		dfs(points, subPoints, i+1, res)
		*subPoints = (*subPoints)[:len(*subPoints)-1]
	}

	return
}

func parollAreaRect(subPoints [][]int) int {
	point1_x := subPoints[0][0]
	point1_y := subPoints[0][1]
	point2_x := subPoints[1][0]
	point2_y := subPoints[1][1]
	point3_x := subPoints[2][0]
	point3_y := subPoints[2][1]
	point4_x := subPoints[3][0]
	point4_y := subPoints[3][1]
	if !((point1_x == point2_x && point3_x == point4_x) ||
		(point1_x == point3_x && point2_x == point4_x) ||
		(point1_x == point4_x && point2_x == point3_x)) == true {
		return 0
	}
	if !((point1_y == point2_y && point3_y == point4_y) ||
		(point1_y == point3_y && point2_y == point4_y) ||
		(point1_y == point4_y && point2_y == point3_y)) == true {
		return 0
	}
	if point1_x == point2_x {
		return int(math.Abs(float64((point2_y - point1_y) * (point3_x - point1_x))))
	}
	if point1_x == point3_x {
		return int(math.Abs(float64((point3_y - point1_y) * (point2_x - point1_x))))
	}
	if point1_x == point4_x {
		return int(math.Abs(float64((point4_y - point1_y) * (point3_x - point1_x))))
	}
	return 0
}

func areaRect(subPoints [][]int) int {
	point1_x := subPoints[0][0]
	point1_y := subPoints[0][1]
	point2_x := subPoints[1][0]
	point2_y := subPoints[1][1]
	point3_x := subPoints[2][0]
	point3_y := subPoints[2][1]
	point4_x := subPoints[3][0]
	point4_y := subPoints[3][1]
	//check if or not three points in one line
	//1, one point
	if (point1_x == point2_x && point2_x == point3_x) ||
		(point1_x == point2_x && point2_x == point4_x) ||
		(point1_x == point3_x && point3_x == point4_x) ||
		(point2_x == point3_x && point3_x == point4_x) {
		return 0
	}
	if (point1_y == point2_y && point2_y == point3_y) ||
		(point1_y == point2_y && point2_y == point4_y) ||
		(point1_y == point3_y && point3_y == point4_y) ||
		(point2_y == point3_y && point3_y == point4_y) {
		return 0
	}
	//2, in one line
	if isBoomerang(subPoints[0], subPoints[1], subPoints[2]) == false ||
		isBoomerang(subPoints[0], subPoints[1], subPoints[3]) == false ||
		isBoomerang(subPoints[1], subPoints[2], subPoints[3]) == false {
		return 0
	}
	// for the parellar side with x and y under the question requirement
	if !((point1_x == point2_x && point3_x == point4_x) ||
		(point1_x == point3_x && point2_x == point4_x) ||
		(point1_x == point4_x && point2_x == point3_x)) == true {
		return 0
	}
	if !((point1_y == point2_y && point3_y == point4_y) ||
		(point1_y == point3_y && point2_y == point4_y) ||
		(point1_y == point4_y && point2_y == point3_y)) == true {
		return 0
	}
	disArray := []int{}
	uniqDis := []int{}
	for i := 0; i < 4; i++ {
		for j := i + 1; j < 4; j++ {
			disArray = append(disArray, dis(subPoints[i][0], subPoints[i][1], subPoints[j][0], subPoints[j][1]))
		}
	}
	// check if or not a rectangle by uniq point distance should be 3 values left
	// and the one value plus other value equal to another value
	sort.Ints(disArray)
	// fmt.Println(disArray)
	for i := 0; i < len(disArray); i++ {
		if i > 0 && disArray[i-1] == disArray[i] {
			continue
		}
		uniqDis = append(uniqDis, disArray[i])
	}
	// fmt.Println(uniqDis)
	if len(uniqDis) == 2 && (uniqDis[0]+uniqDis[0] == uniqDis[1]) {
		fmt.Println(subPoints)
		fmt.Println(uniqDis)
		return uniqDis[0]
	}
	if len(uniqDis) == 3 && (uniqDis[0]+uniqDis[1] == uniqDis[2]) {
		fmt.Println(subPoints)
		fmt.Println(uniqDis)
		return int(math.Sqrt(float64(uniqDis[0])) * math.Sqrt(float64(uniqDis[1])))
	}

	return 0
}

func isBoomerang(point1 []int, point2 []int, point3 []int) bool {
	if point1[0] == point2[0] {
		return point3[0] != point2[0]
	} else {
		k1 := 1.0 * float64((point2[1] - point1[1])) / float64((point2[0] - point1[0]))
		k2 := 1.0 * float64((point3[1] - point1[1])) / float64((point3[0] - point1[0]))
		return k1 != k2
	}
	return false
}

func dis(x int, y int, cx int, cy int) int {
	return (x-cx)*(x-cx) + (y-cy)*(y-cy)
}

// the time limitation could be accepted
func minAreaRect(points [][]int) int {
	if len(points) < 4 {
		return 0
	}
	m := make(map[[2]int]bool)
	min := 0
	const MAX = int(^uint(0) >> 1)
	point1X := MAX
	point1Y := MAX
	point2X := MAX
	point2Y := MAX
	for i := 0; i < len(points); i++ {
		m[[2]int{points[i][0], points[i][1]}] = true
	}
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			a := points[i]
			b := points[j]
			// two conditions are that one is
			//       a
			//  |   /|
			//  |  / |
			//  | /  |
			//  |/   |
			//  b
			// two is
			//  a
			//  |\   |
			//  | \  |
			//  |  \ |
			//  |   \|
			//       b
			// and if it's rectangle, the another two points should be have same x-axis and y-axis
			if (a[0]-b[0])*(a[1]-b[1]) > 0 {
				point1X = int(math.Min(float64(b[0]), float64(a[0])))
				point1Y = int(math.Max(float64(b[1]), float64(a[1])))
				point2X = int(math.Max(float64(b[0]), float64(a[0])))
				point2Y = int(math.Min(float64(b[1]), float64(a[1])))
			} else if (a[0]-b[0])*(a[1]-b[1]) < 0 {
				point1X = int(math.Min(float64(b[0]), float64(a[0])))
				point1Y = int(math.Min(float64(b[1]), float64(a[1])))
				point2X = int(math.Max(float64(b[0]), float64(a[0])))
				point2Y = int(math.Max(float64(b[1]), float64(a[1])))
			}
			if m[[2]int{point1X, point1Y}] == true && m[[2]int{point2X, point2Y}] == true {
				res := int(math.Abs(float64((a[0] - b[0]) * (a[1] - b[1]))))
				// fmt.Println(res)
				if res != 0 {
					if min == 0 {
						min = res
					} else {
						min = int(math.Min(float64(res), float64(min)))
					}
				}

			}
		}
	}
	return min
}