import "math/rand"

// https://leetcode.com/problems/insert-delete-getrandom-o1/
/*
Design a data structure that supports all following operations in average O(1) time.

insert(val): Inserts an item val to the set if not already present.
remove(val): Removes an item val from the set if present.
getRandom: Returns a random element from current set of elements. Each element must have the same probability of being returned.
Example:

// Init an empty set.
RandomizedSet randomSet = new RandomizedSet();

// Inserts 1 to the set. Returns true as 1 was inserted successfully.
randomSet.insert(1);

// Returns false as 2 does not exist in the set.
randomSet.remove(2);

// Inserts 2 to the set, returns true. Set now contains [1,2].
randomSet.insert(2);

// getRandom should return either 1 or 2 randomly.
randomSet.getRandom();

// Removes 1 from the set, returns true. Set now contains [2].
randomSet.remove(1);

// 2 was already in the set, so return false.
randomSet.insert(2);

// Since 2 is the only number in the set, getRandom always return 2.
randomSet.getRandom();

用slice做到random get element
*/
type RandomizedSet struct {
	buff  []int
	cache map[int]int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{
		buff:  make([]int, 0),
		cache: make(map[int]int),
	}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (rs *RandomizedSet) Insert(val int) bool {
	if _, ok := rs.cache[val]; ok {
		return false
	}
	rs.buff = append(rs.buff, val)
	rs.cache[val] = len(rs.buff) - 1
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (rs *RandomizedSet) Remove(val int) bool {
	// 用slice里的最后一个把hash和slice里的要删的值位置覆盖掉 之后然后删除hash里的对应的值和slice最后一个值
	if idx, ok := rs.cache[val]; ok {
		rs.buff[idx] = rs.buff[len(rs.buff)-1]
		rs.cache[rs.buff[len(rs.buff)-1]] = idx
		rs.buff = rs.buff[:len(rs.buff)-1]
		delete(rs.cache, val)
		return true
	}
	return false
}

/** Get a random element from the set. */
func (rs *RandomizedSet) GetRandom() int {
	if len(rs.buff) > 0 {
		return rs.buff[rand.Intn(len(rs.buff))]
	}
	return 0
}
