import "strconv"

// https://leetcode.com/problems/time-based-key-value-store/
/*
Create a timebased key-value store class TimeMap, that supports two operations.

1. set(string key, string value, int timestamp)

Stores the key and value, along with the given timestamp.
2. get(string key, int timestamp)

Returns a value such that set(key, value, timestamp_prev) was called previously, with timestamp_prev <= timestamp.
If there are multiple such values, it returns the one with the largest timestamp_prev.
If there are no values, it returns the empty string ("").


Example 1:

Input: inputs = ["TimeMap","set","get","get","set","get","get"], inputs = [[],["foo","bar",1],["foo",1],["foo",3],["foo","bar2",4],["foo",4],["foo",5]]
Output: [null,null,"bar","bar",null,"bar2","bar2"]
Explanation:
TimeMap kv;
kv.set("foo", "bar", 1); // store the key "foo" and value "bar" along with timestamp = 1
kv.get("foo", 1);  // output "bar"
kv.get("foo", 3); // output "bar" since there is no value corresponding to foo at timestamp 3 and timestamp 2, then the only value is at timestamp 1 ie "bar"
kv.set("foo", "bar2", 4);
kv.get("foo", 4); // output "bar2"
kv.get("foo", 5); //output "bar2"

Example 2:

Input: inputs = ["TimeMap","set","set","get","get","get","get","get"], inputs = [[],["love","high",10],["love","low",20],["love",5],["love",10],["love",15],["love",20],["love",25]]
Output: [null,null,null,"","high","high","low","low"]


Note:

All key/value strings are lowercase.
All key/value strings have length in the range [1, 100]
The timestamps for all TimeMap.set operations are strictly increasing.
1 <= timestamp <= 10^7
TimeMap.set and TimeMap.get functions will be called a total of 120000 times (combined) per test case.
*/

type TimeMap struct {
	m map[string][][]string
}

/** Initialize your data structure here. */
func Constructor() TimeMap {
	m := make(map[string][][]string, 0)
	return TimeMap{m: m}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	time := strconv.Itoa(timestamp)
	this.m[key] = append(this.m[key], []string{time, value})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	time := strconv.Itoa(timestamp)
	if v, ok := this.m[key]; ok {
		index := binarySearch(v, time)
		if index >= 0 {
			return v[index][1]
		}

	}
	return ""
}

func binarySearch(array [][]string, target string) int {
	start := 0
	end := len(array) - 1
	tn, _ := strconv.Atoi(target)
	for (start + 1) < end {
		mid := start + (end-start)/2
		mn, _ := strconv.Atoi(array[mid][0])
		if mn > tn {
			end = mid
		} else if mn < tn {
			start = mid
		} else {
			return mid
		}
	}
	sn, _ := strconv.Atoi(array[start][0])
	en, _ := strconv.Atoi(array[end][0])
	if sn == tn {
		return start
	} else if en == tn {
		return end
	}

	if en < tn {
		return end
	}
	if sn < tn {
		return start
	}
	return -1
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
