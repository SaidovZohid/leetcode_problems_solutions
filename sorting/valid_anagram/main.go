package main

import (
	"fmt"
)

func isAnagram(s string, t string) bool {
	return false
}

func main() {
	s := "anagram"
	t := "nagaram"
	fmt.Println(isAnagram(s, t))
	s = "rat"
	t = "car"
	fmt.Println(isAnagram(s, t))
}
