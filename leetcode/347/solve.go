package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
}

/***
 * Approach: Heap
 * Time: O(n log n)
 * Space: O(n)
 */
type IntHeap []int

var dictionary map[int]int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return dictionary[h[i]] > dictionary[h[j]] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequentHeap(nums []int, k int) []int {
	var ans []int = make([]int, 0)
	dictionary = make(map[int]int)
	items := IntHeap{}
	for _, vl := range nums {
		if _, ok := dictionary[vl]; !ok {
			items = append(items, vl)
		}
		dictionary[vl]++
	}
	heap.Init(&items)
	for k > 0 {
		ans = append(ans, heap.Pop(&items).(int))
		k--
	}
	return ans
}

/***
 * Approach: Sorting, and HashMap
 * Time: O(n log n)
 * Space: O(n)
 */
func topKFrequent(nums []int, k int) []int {
	var keys []int = make([]int, 0)
	var tracker map[int]int = make(map[int]int)
	for _, vl := range nums {
		if _, ok := tracker[vl]; !ok {
			keys = append(keys, vl)
		}
		tracker[vl]++
	}
	sort.Slice(keys, func(i, j int) bool { return tracker[keys[i]] > tracker[keys[j]] })
	return keys[:k]
}
