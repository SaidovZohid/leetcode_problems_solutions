package main

import (
	"fmt"
)

func reverseString(s []byte) []byte {
	if s == nil {
		return s
	}
	l := len(s)

	for i := 0; i < l/2; i++ {
		temp := s[i]
		s[i] = s[l-i-1]
		s[l-i-1] = temp
	}
	return s
}

func main() {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	fmt.Println(string(reverseString(s)))
	s = []byte{'H', 'a', 'n', 'n', 'a', 'h'}
	fmt.Println(string(reverseString(s)))
}
