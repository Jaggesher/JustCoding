package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Case 1:", shortestPathAllKeys([]string{"@.a.#", "###.#", "b.A.B"}))
	fmt.Println("Case 2:", shortestPathAllKeys([]string{"@..aA", "..B#.", "....b"}))
	fmt.Println("Case 3:", shortestPathAllKeys([]string{"@Aa"}))
	fmt.Println("Case 4:", shortestPathAllKeys([]string{"#..#.#.#..#.#.#.....#......#..", ".#.......#....#A.....#.#......", "#....#.....#.........#........", "...#.#.........#..@....#....#.", ".#.#.##...#.........##....#..#", "..........#..#..###....##..#.#", ".......#......#...#...#.....c#", ".#...#.##......#...#.###...#..", "..........##...#.......#......", "#...#.........a#....#.#.##....", "..#..#...#...#..#....#.....##.", "..........#...#.##............", "...#....#..#.........#..D.....", "....#E.#....##................", "...........##.#.......#.#....#", "...#..#...#.#............#e...", "..#####....#.#...........##..#", "##......##......#.#...#..#.#..", ".#F.......#..##.......#....#..", "............#....#..#..#...#..", ".............#...#f...#..##...", "....#..#...##.........#..#..#.", ".....#.....##.###..##.#......#", ".#..#.#...#.....#........###..", ".....#.#...#...#.....#.....#..", "##.....#....B.....#..#b.......", ".####....##..#.##..d.#......#.", "..#.....#....##........##...##", "...#...#...C..#..#....#.......", "#.....##.....#.#......#......."}))
	fmt.Println("Case 5:", shortestPathAllKeys([]string{"@abcdeABCDEFf"}))
}

type Node struct {
	x     int
	y     int
	keys  string
	count int
}

func shortestPathAllKeys(grid []string) int {
	var tracker [30][30]map[string]int
	for i := 0; i < 30; i++ {
		for j := 0; j < 30; j++ {
			tracker[i][j] = make(map[string]int)
		}
	}

	m, n := len(grid), len(grid[0])
	keys, stX, stY := getKeysAndPoint(grid)
	if keys == 0 {
		return 0
	}

	var queue []Node = make([]Node, 0)
	var dir [][]int = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	queue = append(queue, Node{x: stX, y: stY, keys: "", count: 0})
	tracker[stX][stY][""] = 1
	for len(queue) > 0 {
		tm := queue[0]
		queue = queue[1:]
		for _, vl := range dir {
			x, y, count := tm.x+vl[0], tm.y+vl[1], tm.count+1
			if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] != '#' {
				flag := true
				newKey := tm.keys

				if grid[x][y] >= 'A' && grid[x][y] <= 'Z' {
					flag = isKeyPresent(tm.keys, rune('a'+(grid[x][y]-'A')))
				} else if grid[x][y] >= 'a' && grid[x][y] <= 'z' && !isKeyPresent(tm.keys, rune(grid[x][y])) {
					temp := []byte(newKey)
					temp = append(temp, grid[x][y])
					sort.Slice(temp, func(i, j int) bool { return temp[i] < temp[j] })
					newKey = string(temp)
					if len(newKey) == keys {
						return count
					}
				}

				if tracker[x][y][newKey] > 0 {
					flag = false
				}

				if flag {
					queue = append(queue, Node{x: x, y: y, keys: newKey, count: count})
					tracker[x][y][newKey] = count
				}
			}
		}

	}
	return -1
}

func isKeyPresent(keys string, key rune) bool {
	for _, vl := range keys {
		if vl == key {
			return true
		}
	}
	return false
}

func getKeysAndPoint(grid []string) (keys int, x int, y int) {
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '@' {
				x, y = i, j
			} else if grid[i][j] >= 'a' && grid[i][j] <= 'z' {
				keys++
			}
		}
	}
	return keys, x, y
}
