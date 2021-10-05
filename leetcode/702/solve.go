package main

import (
	"fmt"
)

func main() {
	fmt.Println("Bo test case")
}

type ArrayReader struct {
}

func (this *ArrayReader) get(index int) int { return -1 } // not implemented

func search(reader ArrayReader, target int) int {
	left, right := 0, 1
	for reader.get(right) < target {
		left = right
		right <<= 1
	}
	for left <= right {
		mid := left + ((right - left) >> 1)
		tm := reader.get(mid)
		if target == tm {
			return mid
		} else if target > tm {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
