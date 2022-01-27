package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println("Case 1: ", kthSmallest([][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}, 8))
	fmt.Println("Case 2: ", kthSmallest([][]int{{-5}}, 1))
	fmt.Println("Case 3: ", kthSmallest([][]int{{1, 2}, {1, 3}}, 3))
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kthSmallest(matrix [][]int, k int) int {
	var hp IntHeap = make(IntHeap, k)
	heap.Init(&hp)
	count := 0
	for _, first := range matrix {
		for _, sec := range first {
			if count < k {
				hp[count] = sec
			} else if hp[0] > sec {
				heap.Pop(&hp)
				heap.Push(&hp, sec)
			}
			if count+1 == k {
				heap.Init(&hp)
			}
			count++
		}
	}

	return heap.Pop(&hp).(int)
}
