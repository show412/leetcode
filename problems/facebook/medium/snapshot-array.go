import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/snapshot-array/
/*
Implement a SnapshotArray that supports the following interface:

SnapshotArray(int length) initializes an array-like data structure with the given length.  Initially, each element equals 0.
void set(index, val) sets the element at the given index to be equal to val.
int snap() takes a snapshot of the array and returns the snap_id: the total number of times we called snap() minus 1.
int get(index, snap_id) returns the value at the given index, at the time we took the snapshot with the given snap_id


Example 1:

Input: ["SnapshotArray","set","snap","set","get"]
[[3],[0,5],[],[0,6],[0,0]]
Output: [null,null,0,null,5]
Explanation:
SnapshotArray snapshotArr = new SnapshotArray(3); // set the length to be 3
snapshotArr.set(0,5);  // Set array[0] = 5
snapshotArr.snap();  // Take a snapshot, return snap_id = 0
snapshotArr.set(0,6);
snapshotArr.get(0,0);  // Get the value of array[0] with snap_id = 0, return 5


Constraints:

1 <= length <= 50000
At most 50000 calls will be made to set, snap, and get.
0 <= index < length
0 <= snap_id < (the total number of times we call snap())
0 <= val <= 10^9
*/
// 这道题的关键是定义数据结构 map里面的value也是map
/*
The case will be LTE
因为在get的时候先排序 又循环 是O(nlogn + n)
https://leetcode.com/submissions/detail/267395654/testcase/#
*/
type SnapshotArray struct {
	A      map[int]map[int]int
	snapID int
}

func Constructor(length int) SnapshotArray {
	arr := make(map[int]map[int]int, 0)
	for i := 0; i < length; i++ {
		arr[i] = make(map[int]int, 0)
		arr[i][0] = 0
	}
	return SnapshotArray{arr, 0}
}

func (this *SnapshotArray) Set(index int, val int) {
	this.A[index][this.snapID] = val
	return
}

func (this *SnapshotArray) Snap() int {
	id := this.snapID
	this.snapID++
	return id
}

func (this *SnapshotArray) Get(index int, snap_id int) int {
	fmt.Println("*****")
	fmt.Println(this.A)
	j := 0
	keys := make([]int, len(this.A[index]))
	for k := range this.A[index] {
		keys[j] = k
		j++
	}
	sort.Ints(keys)
	fmt.Println(keys)
	for i := len(keys) - 1; i >= 0; i-- {
		if keys[i] <= snap_id {
			return this.A[index][keys[i]]
		}
	}
	return 0
	// i := sort.Search(len(keys), func(i int) bool { return keys[i] >= snap_id })
	// fmt.Println(i)
	// if i > len(keys)-1 {
	// 	i = len(keys) - 1
	// }
	// if keys[i] > snap_id && i > 0 {
	// 	i = i - 1
	// }
	// return this.A[index][keys[i]]
}

/**
 * Your SnapshotArray object will be instantiated and called as such:
 * obj := Constructor(length);
 * obj.Set(index,val);
 * param_2 := obj.Snap();
 * param_3 := obj.Get(index,snap_id);
 */

// 这个方式就比较好 get 只是O(n)的复杂度
// SnapshotArray can snap a array
type SnapshotArray struct {
	snaps   []map[int]int
	current map[int]int
	id      int
}

// Constructor resturn s a SnapshotArray
func Constructor(length int) SnapshotArray {
	s := make([]map[int]int, 0, 128)
	c := make(map[int]int, 32)
	return SnapshotArray{
		snaps:   s,
		current: c,
		id:      0,
	}
}

// Set val in index
func (sa *SnapshotArray) Set(index int, val int) {
	sa.current[index] = val
}

// Snap make snapshot
func (sa *SnapshotArray) Snap() int {
	sa.snaps = append(sa.snaps, sa.current)
	sa.current = make(map[int]int)
	sa.id++
	return sa.id - 1
}

// Get returns val in the snap
func (sa *SnapshotArray) Get(index int, snap int) int {
	for id := snap; id >= 0; id-- {
		history := sa.snaps[id]
		if t, ok := history[index]; ok {
			return t
		}
	}
	return 0 //default value
}
