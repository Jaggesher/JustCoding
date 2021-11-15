package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println("Case 1:", minCost([][]int{{1, 1, 1, 1}, {2, 2, 2, 2}, {1, 1, 1, 1}, {2, 2, 2, 2}}))
	fmt.Println("Case 2:", minCost([][]int{{1, 1, 3}, {3, 2, 2}, {1, 1, 4}}))
	fmt.Println("Case 3:", minCost([][]int{{1, 2}, {4, 3}}))
}

func minCost(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	var tracker [101][101]int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			tracker[i][j] = -1
		}
	}

	var dirs [][]int = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	pq := make(PriorityQueue, 0)
	heap.Push(&pq, &Cell{i: 0, j: 0, weight: 0})
	tracker[0][0] = 0
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Cell)
		if tracker[item.i][item.j] >= 0 && tracker[item.i][item.j] < item.weight {
			continue
		}
		for index, dir := range dirs {
			temp := Cell{i: item.i + dir[0], j: item.j + dir[1], weight: item.weight}
			if temp.i < 0 || temp.j < 0 || temp.i >= m || temp.j >= n {
				continue
			}
			if (index + 1) != grid[item.i][item.j] {
				temp.weight++
			}

			if tracker[temp.i][temp.j] == -1 || tracker[temp.i][temp.j] > temp.weight {
				tracker[temp.i][temp.j] = temp.weight
				heap.Push(&pq, &temp)
			}
		}
	}

	return tracker[m-1][n-1]
}

// Priority Queue
type Cell struct {
	i, j, weight int
}

type PriorityQueue []*Cell

func (pq PriorityQueue) Len() int            { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool  { return pq[i].weight < pq[j].weight }
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(*Cell)) }

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}
