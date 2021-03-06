import "sort"

// https://leetcode.com/problems/accounts-merge/
/*
Given a list accounts, each element accounts[i] is a list of strings,
where the first element accounts[i][0] is a name,
and the rest of the elements are emails representing emails of the account.

Now, we would like to merge these accounts.
Two accounts definitely belong to the same person
if there is some email that is common to both accounts.
Note that even if two accounts have the same name,
they may belong to different people as people could have the same name.
A person can have any number of accounts initially,
but all of their accounts definitely have the same name.

After merging the accounts,
return the accounts in the following format:
the first element of each account is the name,
and the rest of the elements are emails in sorted order.
The accounts themselves can be returned in any order.

Example 1:
Input:
accounts = [["John", "johnsmith@mail.com", "john00@mail.com"],
["John", "johnnybravo@mail.com"],
["John", "johnsmith@mail.com", "john_newyork@mail.com"],
["Mary", "mary@mail.com"]]
Output: [["John", 'john00@mail.com', 'john_newyork@mail.com', 'johnsmith@mail.com'],
["John", "johnnybravo@mail.com"], ["Mary", "mary@mail.com"]]
Explanation:
The first and third John's are the same person as they have the common email "johnsmith@mail.com".
The second John and Mary are different people as none of their email addresses are used by other accounts.
We could return these lists in any order, for example the answer [['Mary', 'mary@mail.com'], ['John', 'johnnybravo@mail.com'],
['John', 'john00@mail.com', 'john_newyork@mail.com', 'johnsmith@mail.com']] would still be accepted.
Note:

The length of accounts will be in the range [1, 1000].
The length of accounts[i] will be in the range [1, 10].
The length of accounts[i][j] will be in the range [1, 30].
*/
// find union solution
type DSU struct {
	parent []int
}

func constructor() DSU {
	parent := make([]int, 10000)
	for k := range parent {
		parent[k] = k
	}
	return DSU{parent: parent}
}

func (ds DSU) find(x int) int {
	if ds.parent[x] != x {
		ds.parent[x] = ds.find(ds.parent[x])
	}
	return ds.parent[x]
}

func (ds DSU) union(x, y int) {
	ds.parent[ds.find(x)] = ds.find(y)
}

func accountsMerge(accounts [][]string) [][]string {
	ds := constructor()
	emailToName := make(map[string]string)
	emailToID := make(map[string]int)
	id := 0
	for _, account := range accounts {
		name := ""
		for _, v := range account {
			if name == "" {
				name = v
				continue
			}
			emailToName[v] = name
			if _, ok := emailToID[v]; !ok {
				emailToID[v] = id
				id++
			}
			// 处理这些ID的父子关系
			ds.union(emailToID[account[1]], emailToID[v])
		}
	}
	// emailToID 记录了 email对ID的映射
	ans := make(map[int][]string)
	for k := range emailToName {
		// 因为这些ID的父子关系 把一样父id的email放到ans的一样的hash里
		index := ds.find(emailToID[k])
		ans[index] = append(ans[index], k)
	}
	var result [][]string
	var arr []string
	for _, v := range ans {
		sort.Strings(v)
		arr = nil
		name := emailToName[v[0]]
		arr = append(arr, name)
		arr = append(arr, v...)
		result = append(result, arr)
	}
	return result
}
