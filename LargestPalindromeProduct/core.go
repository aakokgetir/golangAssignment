package LargestPalindromeProduct

import (
	"strconv"
)

func ReverseStr(s string) string {
	rev := ""
	for i := len(s) - 1; i >= 0; i-- {
		rev += string(s[i])
	}
	return rev
}

func IsPalindrome(n int) bool {
	s := strconv.Itoa(n)
	return s == ReverseStr(s)
}

func LargestNumber(digits int) int {
	s := ""
	for i := 0; i < digits; i++ {
		s += "9"
	}
	val, _ := strconv.Atoi(s)
	return val
}

func SmallestNumber(digits int) int {
	s := "1"
	for i := 0; i < digits-1; i++ {
		s += "0"
	}
	val, _ := strconv.Atoi(s)
	return val
}

func LargestPalindromeNumbersForProduct(digits int) (int, int, int) {
	mx := LargestNumber(digits)
	sm := SmallestNumber(digits)
	largestProduct, p1, p2 := 0, 0, 0
	for i := mx; i >= sm; i-- {
		for j := mx; j >= sm; j-- {
			product := i * j
			if IsPalindrome(product) {
				if product > largestProduct {
					largestProduct, p1, p2 = product, i, j
				}
			}
		}
	}
	return p1, p2, largestProduct
}
