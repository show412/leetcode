// https://leetcode.com/problems/lfu-cache/
/*
Design and implement a data structure for Least Frequently Used (LFU) cache. It should support the following operations: get and put.

get(key) - Get the value (will always be positive) of the key if the key exists in the cache, otherwise return -1.
put(key, value) - Set or insert the value if the key is not already present. When the cache reaches its capacity, it should invalidate the least frequently used item before inserting a new item. For the purpose of this problem, when there is a tie (i.e., two or more keys that have the same frequency), the least recently used key would be evicted.

Note that the number of times an item is used is the number of calls to the get and put functions for that item since it was inserted. This number is set to zero when the item is removed.



Follow up:
Could you do both operations in O(1) time complexity?



Example:

LFUCache cache = new LFUCache( 2 /* capacity */ );
/*
cache.put(1, 1);
cache.put(2, 2);
cache.get(1);       // returns 1
cache.put(3, 3);    // evicts key 2
cache.get(2);       // returns -1 (not found)
cache.get(3);       // returns 3.
cache.put(4, 4);    // evicts key 1.
cache.get(1);       // returns -1 (not found)
cache.get(3);       // returns 3
cache.get(4);       // returns 4
和LRU的区别是 LRU是最近最少使用淘汰，LFU是在使用次数最少的前提下淘汰最早的那个页面
LRU可以用双向链表实现 LFU也是用双向链表实现 但是多了一个hash记录访问次数对应的node
存储的是所有访问次数为i的最后一个节点。接下来就简单了，只要某个节点被访问，
就令它的次数num加1，从原来的位置摘除，并插入到次数为num的地方。
如果这里已经有次数为num 的节点，则插入到mcount【num】这地址的后面，
如果不存在，就插入到原来的mcount【num-1】的后面，成为第一个次数为num的节点
https://blog.csdn.net/qq_41445952/article/details/81878065
*/
class LFUCache {

    private class Node{
        int key;
        int val;
        int count;
        Node next;
        Node pre;
        Node(int key, int val){this.key = key; this.val = val;this.count = 1;}

    }

    private HashMap<Integer, Node> hashMap;
    private HashMap<Integer, Node> headMap;
    private int leastCount = 1;
    private int capacity;

    public LFUCache(int capacity) {
        hashMap = new HashMap<>();
        headMap = new HashMap<>();
        this.capacity = capacity;
    }

    public int get(int key) {
        Node node = hashMap.get(key);
        if(node == null) return -1;
        update(node);
        return node.val;
    }

    public void put(int key, int value) {
        if(capacity == 0) return;
        Node node = hashMap.get(key);
        if(node == null){
            node = new Node(key, value);
            if(hashMap.size() ==  capacity){
                removeLeast();
            }
            hashMap.put(key, node);
            insert(node);
            leastCount = 1;
        }else{
            node.val = value;
            update(node);
        }
    }

    public void insert(Node node){
        Node tmp = headMap.get(node.count);
        if(tmp == null){
            tmp = new Node(0,0); // an extra head.
            tmp.next = tmp;
            tmp.pre = tmp;
            headMap.put(node.count, tmp);
        }
        node.next = tmp.next;
        node.next.pre = node;
        node.pre = tmp;
        tmp.next = node;
    }

    public void update(Node node){
        remove(node);
        node.count++;
        insert(node);
    }

    public boolean remove(Node node){
        if(node.next == node){
            return false;
        }else{
            if(node.count == leastCount && node.pre == node.next){
                // last one in such count;
                leastCount++;
            }
            Node preNode = node.pre;
            preNode.next = node.next;
            node.next.pre = preNode;
            return true;
        }
    }


    public void removeLeast(){
        Node lastNode = headMap.get(leastCount).pre;
        remove(lastNode);
        hashMap.remove(lastNode.key);
    }
}

// 用三个hash来存储 一个是key value的hash  一个是key count的， 一个是count对应的一个link
class LFUCache {
    private HashMap<Integer,Integer> vals;
    private HashMap<Integer,Integer> counter;
    private HashMap<Integer,LinkedHashSet<Integer>> list;
    private int min = -1;
    private int capacity;
    public LFUCache(int capacity) {
        this.capacity = capacity;
        vals = new HashMap<>();
        counter=new HashMap<>();
        list = new HashMap<>();
        list.put(1,new LinkedHashSet<>());
    }

    public int get(int key) {
        if(!vals.containsKey(key)){
            return -1;
        }
        int count = counter.get(key);
        counter.put(key,count+1);
        list.get(count).remove(key);
        if(count == min && list.get(count).size()==0)
            min++;
        if(!list.containsKey(count+1))
            list.put(count+1,new LinkedHashSet<>());
        list.get(count+1).add(key);
        return vals.get(key);
    }

    public void put(int key, int value) {
        if(capacity<=0)
            return;
        if(vals.containsKey(key)){
            vals.put(key,value);
            get(key);
            return;
        }else {
            if(vals.size()>=capacity){
                int evict = list.get(min).iterator().next();
                list.get(min).remove(evict);
                vals.remove(evict);
                counter.remove(evict);
            }
            min =1;
            vals.put(key,value);
            counter.put(key,1);
            list.get(1).add(key);
        }
    }
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * LFUCache obj = new LFUCache(capacity);
 * int param_1 = obj.get(key);
 * obj.put(key,value);
 */
