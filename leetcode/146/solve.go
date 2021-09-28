package main

import "fmt"

type DList struct {
	Val  int
	key  int
	Next *DList
	Prev *DList
}

type LRUCache struct {
	capacity int
	head     *DList
	tail     *DList
	mp       map[int]*DList
}

func Constructor(capacity int) LRUCache {
	obj := new(LRUCache)
	obj.capacity = capacity
	obj.mp = make(map[int]*DList)
	obj.head = &DList{Val: 0, key: 0, Next: nil, Prev: nil}
	obj.tail = &DList{Val: 0, key: 0, Next: nil, Prev: nil}
	obj.head.Next = obj.tail
	obj.tail.Prev = obj.head
	return *obj
}

func (this *LRUCache) Get(key int) int {
	vl, present := this.mp[key]
	if present {
		this.remove(vl)
		this.add(vl)
		return vl.Val
	}
	return -1
}

func (this *LRUCache) remove(node *DList) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (this *LRUCache) add(node *DList) {
	next := this.head.Next
	this.head.Next = node
	node.Prev = this.head
	node.Next = next
	next.Prev = node
}

func (this *LRUCache) Put(key int, value int) {
	vl, present := this.mp[key]
	if present {
		vl.Val = value
		this.remove(vl)
		this.add(vl)
	} else if this.capacity > len(this.mp) {
		node := DList{Val: value, key: key, Next: nil, Prev: nil}
		this.add(&node)
		this.mp[key] = &node
	} else {
		node := this.tail.Prev
		delete(this.mp, node.key)
		node.Val = value
		node.key = key
		this.remove(node)
		this.add(node)
		this.mp[key] = node
	}
}

func main() {
	fmt.Println("No test case")
	/**
	 * Your LRUCache object will be instantiated and called as such:
	 * obj := Constructor(capacity);
	 * param_1 := obj.Get(key);
	 * obj.Put(key,value);
	 */
}
