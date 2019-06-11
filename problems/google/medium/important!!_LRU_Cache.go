// https://leetcode.com/problems/lru-cache/
// use hash to store the key value
// use double linklist to implement the remove for least recently
type LRUCache struct {
	hashMap  map[int]*LinkList // 这里用指针作为value 否则后面取出来的时候是值传递会导致不能改变linklist
	capacity int
	head     *LinkList
	tail     *LinkList
}

// for the remove for least recently
type LinkList struct {
	key   int
	value int
	pre   *LinkList // 结构体的递归定义用指针
	next  *LinkList // 结构体的递归定义用指针
}

func Constructor(capacity int) LRUCache {
	m := make(map[int]*LinkList, capacity)
	h := LinkList{key: -1, value: -1, pre: nil, next: nil}
	t := LinkList{key: -1, value: -1, pre: nil, next: nil}
	t.pre = &h
	h.next = &t
	return LRUCache{hashMap: m, capacity: capacity, head: &h, tail: &t}
}

// 在go中 如果是方法，则接收者定义为对象还是对象指针，都可以接收对象和对象指针的调用
func (this *LRUCache) Get(key int) int {
	if v, ok := this.hashMap[key]; ok {
		/* 如果这里hashmap里存的不是指针 而是一个linkList
		这个v其实是原有的一个复制的值 导致原有的那个node并不会脱离开head.next
		因为随着值传递复制 head.next下面其实是指了两个1 后面的一系列操作其实只是删了一个1
		*/
		// link the pre of v to the next of v
		v.pre.next = v.next
		v.next.pre = v.pre
		// move the v to the bottom of linklist
		v.pre = this.tail.pre
		this.tail.pre = v
		v.pre.next = v
		v.next = this.tail

		// v.pre = this.tail.pre
		// v.next = this.tail
		// this.tail.pre.next = &v
		// this.tail.pre = &v
		return v.value
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	// 这句逻辑要注意 因为可能会对同一个key进行update
	// 即使没有update put一样的value也算访问一次 要把这个node放到link最后
	if this.Get(key) != -1 {
		this.hashMap[key].value = value
		return
	}
	if len(this.hashMap) == this.capacity {
		// delete the value from hash
		delete(this.hashMap, this.head.next.key)

		// delete the first node from linklist
		this.head.next = this.head.next.next
		this.head.next.pre = this.head
	}
	// move the node to the pre of tail
	insert := LinkList{key: key, value: value, pre: this.tail.pre, next: this.tail}

	this.tail.pre.next = &insert
	this.tail.pre = &insert
	// insert the node as value of hash

	this.hashMap[key] = &insert

}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
