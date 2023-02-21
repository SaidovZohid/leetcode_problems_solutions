package main

import (
	"fmt"
	"unicode"
)

func isPalindrome(s string) bool {
	s = remove(s)
	l := len(s)
	for i := 0; i < l/2; i++ {
		if s[i] != s[l-i-1] {
			return false
		}
	}
	return true
}

func remove(s string) string {
	var newS string
	for _, v := range s {
		if unicode.IsLetter(v) {
			newS += string(unicode.ToLower(v))
		} else if unicode.IsNumber(v) {
			newS += string(v)
		}
	}

	return newS
}

func main() {
	s := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(s))
	s = "race a car"
	fmt.Println(isPalindrome(s))
	s = "0P"
	fmt.Println(isPalindrome(s))
}
