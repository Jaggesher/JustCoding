package main

import "fmt"

func main() {
	fmt.Println(getSum(0, 0))
	fmt.Println(getSum(-1, 0))
	fmt.Println(getSum(0, -1))
	fmt.Println(getSum(-1000, -1000))
	fmt.Println(getSum(1000, 1000))
}

func getSum(a int, b int) int {
	if a >= 0 && b >= 0 {
		return sum(a, b)
	}
	if a < 0 && b < 0 {
		return -sum(abs(a), abs(b))
	}
	if a < 0 {
		if abs(a) > b {
			return -diff(abs(a), b)
		}
		return diff(a, b)
	}
	if abs(b) > a {
		return -diff(abs(b), a)
	}
	return diff(abs(b), a)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func sum(a int, b int) int {
	sum, extra := 0, 0
	for a > 0 || b > 0 {
		extra += (a % 10) + (b % 10)
		sum *= 10
		sum += extra % 10
		extra /= 10
		a /= 10
		b /= 10
	}

	for sum > 0 {
		extra *= 10
		extra += (sum % 10)
		sum /= 10
	}
	return extra
}

func diff(a int, b int) int {
	if a < b {
		a, b = b, a
	}
	diff, borrowed := 0, 0
	for a > 0 {
		la, lb := a%10, b%10
		lb += borrowed
		borrowed = 0
		if la < lb {
			la += 10
			borrowed = 1
		}
		diff *= 10
		diff += (la - lb)
		a /= 10
		b /= 10
	}
	tm := 0
	for diff > 0 {
		tm *= (diff % 10)
		diff /= 10
	}
	return tm
}
