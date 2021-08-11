package main

import (
	"fmt"
)

func main() {
	fmt.Println("Case 1:", canIWin(10, 11))
	fmt.Println("Case 2:", canIWin(10, 0))
	fmt.Println("Case 3:", canIWin(10, 1))
	fmt.Println("Case 4:", canIWin(20, 200))
	fmt.Println("Case 5:", canIWin(4, 6))
}

func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	var cache []int = make([]int, 1<<maxChoosableInteger)
	var mask int = 0
	for i := 0; i < maxChoosableInteger; i++ {
		mask += (1 << i)
	}
	if desiredTotal > int(((maxChoosableInteger+1)*maxChoosableInteger)/2) {
		return false
	}
	return PlayOptimally(desiredTotal, maxChoosableInteger, mask, cache) == 1
}

func PlayOptimally(currentDesire int, maxChoosableInteger int, mask int, cache []int) int {
	if cache[mask] != 0 {
		return cache[mask]
	}
	for i := 0; i < maxChoosableInteger; i++ {
		tm := mask & (1 << i)
		if tm > 0 {
			rest := currentDesire - (i + 1)
			if rest <= 0 {
				cache[mask] = 1
				return 1
			} else if PlayOptimally(rest, maxChoosableInteger, mask-(1<<i), cache) != 1 {
				cache[mask] = 1
				return 1
			}
		}
	}
	cache[mask] = -1
	return -1
}
