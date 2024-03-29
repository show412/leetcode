/*
 * @Author: hongwei.sun
 * @Date: 2021-01-22 18:45:51
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-03-29 23:23:06
 * @Description: file content
 */
 # // https://leetcode.com/problems/lru-cache/
 # /*
 # Design and implement a data structure for Least Recently Used (LRU) cache.
 # It should support the following operations: get and put.
 
 # get(key) - Get the value (will always be positive) of the key
 # if the key exists in the cache, otherwise return -1.
 # put(key, value) - Set or insert the value if the key is not already present.
 # When the cache reached its capacity,
 # it should invalidate the least recently used item before inserting a new item.
 
 # The cache is initialized with a positive capacity.
 
 # Follow up:
 # Could you do both operations in O(1) time complexity?
 
 # Example:
 
 # LRUCache cache = new LRUCache( 2 capacity );
 
 # cache.put(1, 1);
 # cache.put(2, 2);
 # cache.get(1);       // returns 1
 # cache.put(3, 3);    // evicts key 2
 # cache.get(2);       // returns -1 (not found)
 # cache.put(4, 4);    // evicts key 1
 # cache.get(1);       // returns -1 (not found)
 # cache.get(3);       // returns 3
 # cache.get(4);       // returns 4
 
 # */
 /*
 LRU is about how to design one data structure
 1, one double link node,   
 type node struct {
	 Val [2]int{key, value}
	 pre *node
	 next *node
 }
 2, hashmap,  key => input key,  value => pointer to node
 3, Left node point to least recently use node
 Right node point to latest recently use node
 
 */
 
 type Node struct {
	 Key int
	 Value int
	 Pre *Node
	 Next *Node
 }
 
 type LRUCache struct {
	 Capacity int
	 Cache map[int]*Node
	 Left *Node
	 Right *Node
 }
 
 
 func Constructor(capacity int) LRUCache {
	 cache := make(map[int]*Node,capacity)
	 left := &Node{}
	 right := &Node{}
	 left.Next = right
	 right.Pre = left
	 return LRUCache {
		 Capacity : capacity,
		 Cache : cache,
		 Left : left,
		 Right : right, 
	 }
 }
 
 func (this *LRUCache) remove(node *Node) {
	 pre, next := node.Pre, node.Next
	 pre.Next = next
	 next.Pre = pre
	 // notice we need to delete cache map otherwise it's still exists
	 delete(this.Cache, node.Key)
 }
 
 func (this *LRUCache) insert(node *Node) {
	 pre := this.Right.Pre 
	 pre.Next = node
	 node.Pre = pre
	 node.Next = this.Right
	 this.Right.Pre = node
	 // need to put node into cache map
	 this.Cache[node.Key] = node
 }
 func (this *LRUCache) Get(key int) int {
	 if v, ok := this.Cache[key]; ok {
		 this.remove(v)
		 this.insert(v)
		 return v.Value
	 }
	 return -1
 }
 
 
 func (this *LRUCache) Put(key int, value int)  {
	 if v, ok := this.Cache[key]; ok {
		 this.remove(v)
	 }
	 this.Cache[key] = &Node{Key:key, Value:value}
	 this.insert(this.Cache[key])
	 if(len(this.Cache) > this.Capacity) {
		 least := this.Left.Next
		 this.remove(least)
	 }
 }
 
 
 /**
  * Your LRUCache object will be instantiated and called as such:
  * obj := Constructor(capacity);
  * param_1 := obj.Get(key);
  * obj.Put(key,value);
  */