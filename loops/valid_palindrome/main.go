package main

import (
	"fmt"
)

func isPalindrome(s string) bool {
	ns := make([]int32, 0, len(s))
	for _, v := range s {
		if v >= 65 && v <= 90 {
			ns = append(ns, v+32)
		} else if (v >= 97 && v <= 122) || (v >= 48 && v <= 57) {
			ns = append(ns, v)
		}
	}
	for i := 0; i < len(ns)/2; i++ {
		if ns[i] != ns[len(ns)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	s := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(s))
	s = "race a car"
	fmt.Println(isPalindrome(s))
	s = "0P"
	fmt.Println(isPalindrome(s))
}
