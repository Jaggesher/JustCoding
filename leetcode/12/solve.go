package main

import "fmt"

func main() {
	fmt.Println("Case 1:", intToRoman(3))
	fmt.Println("Case 2:", intToRoman(4))
	fmt.Println("Case 3:", intToRoman(9))
	fmt.Println("Case 4:", intToRoman(58))
	fmt.Println("Case 5:", intToRoman(1994))
}

type NumMap struct {
	val   int
	roman string
}

func intToRoman(num int) string {
	var ans string = ""
	var mapper []NumMap = []NumMap{{val: 1000, roman: "M"}, {val: 900, roman: "CM"},
		{val: 500, roman: "D"}, {val: 400, roman: "CD"}, {val: 100, roman: "C"},
		{val: 90, roman: "XC"}, {val: 50, roman: "L"}, {val: 40, roman: "XL"},
		{val: 10, roman: "X"}, {val: 9, roman: "IX"}, {val: 5, roman: "V"},
		{val: 4, roman: "IV"}, {val: 1, roman: "I"}}
	for _, el := range mapper {
		for num >= el.val {
			ans += el.roman
			num -= el.val
		}
	}
	return ans
}
