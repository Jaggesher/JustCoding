package main

import (
	"fmt"
)

func main() {
	fmt.Println("No Test case")
}

// This is the robot's control interface.
// You should not implement it, or speculate about its implementation
type Robot struct {
}

// Returns true if the cell in front is open and robot moves into the cell.
// Returns false if the cell in front is blocked and robot stays in the current cell.
func (robot *Robot) Move() bool { return true }

// Robot will stay in the same cell after calling TurnLeft/TurnRight.
// Each turn will be 90 degrees.
func (robot *Robot) TurnLeft()  {}
func (robot *Robot) TurnRight() {}

// Clean the current cell.
func (robot *Robot) Clean() {}

type Cell struct {
	x int
	y int
}

func cleanRoom(robot *Robot) {
	var tracker map[Cell]bool = make(map[Cell]bool)
	var dirs [][]int = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	solve(Cell{x: 0, y: 0}, 0, robot, tracker, dirs)
}

func solve(cell Cell, dir int, robot *Robot, tracker map[Cell]bool, dirs [][]int) int {
	robot.Clean()
	tracker[cell] = true
	currentFace := dir
	for i, vl := range dirs {
		changeDirection(robot, currentFace, i)
		nextCell := Cell{x: cell.x + vl[0], y: cell.y + vl[1]}
		if tracker[nextCell] == false && robot.Move() {
			currentFace = solve(nextCell, i, robot, tracker, dirs)
		} else {
			currentFace = i
		}
	}
	changeDirection(robot, currentFace, (dir+2)%4)
	robot.Move()
	return (dir + 2) % 4
}

func changeDirection(robot *Robot, current, next int) {
	if current == next {
		return
	}
	if ((current + 1) % 4) == next {
		robot.TurnRight()
	} else if ((current + 3) % 4) == next {
		robot.TurnLeft()
	} else {
		robot.TurnRight()
		robot.TurnRight()
	}
}
