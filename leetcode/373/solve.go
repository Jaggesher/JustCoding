package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println("Case 1: ", kSmallestPairs([]int{1, 7, 11}, []int{2, 4, 6}, 3))
	fmt.Println("Case 2: ", kSmallestPairs([]int{1, 1, 2}, []int{1, 2, 3}, 2))
	fmt.Println("Case 3: ", kSmallestPairs([]int{1, 2}, []int{3}, 3))
}

/***
 * Approach: 1
 * Time: O(k * n)
 * Space: O(n)
 */

func kSmallestPairs1(nums1 []int, nums2 []int, k int) [][]int {
	if k > len(nums1)*len(nums2) {
		k = len(nums1) * len(nums2)
	}
	var ans [][]int = make([][]int, k)
	var tracker []int = make([]int, len(nums1))
	for index := 0; index < k; index++ {
		min, min_index := 0, -1
		for i, vl := range nums1 {
			if tracker[i] < len(nums2) && (min_index < 0 || (nums2[tracker[i]]+vl) < min) {
				min, min_index = (nums2[tracker[i]] + vl), i
			}
		}
		ans[index] = []int{nums1[min_index], nums2[tracker[min_index]]}
		tracker[min_index]++
	}
	return ans
}

/***
 * Approach: 2
 * Time: O(k log n)
 * Space: O(n)
 */

type pair struct {
	i, j, sum int
}

type pairHeap []*pair

func (h pairHeap) Len() int            { return len(h) }
func (h pairHeap) Less(i, j int) bool  { return h[i].sum < h[j].sum }
func (h pairHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *pairHeap) Push(x interface{}) { *h = append(*h, x.(*pair)) }
func (h *pairHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*h = old[0 : n-1]
	return item
}
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	var ans [][]int = make([][]int, 0, k)
	pq := make(pairHeap, len(nums1))
	for index := range nums1 {
		pq[index] = &pair{i: index, j: 0, sum: nums1[index] + nums2[0]}
	}
	heap.Init(&pq)
	for pq.Len() > 0 && k > 0 {
		item := heap.Pop(&pq).(*pair)
		ans = append(ans, []int{nums1[item.i], nums2[item.j]})
		k--
		if (item.j + 1) == len(nums2) {
			continue
		}
		heap.Push(&pq, &pair{i: item.i, j: item.j + 1, sum: nums1[item.i] + nums2[item.j+1]})
	}
	return ans
}
