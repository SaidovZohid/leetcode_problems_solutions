package main

import (
	"fmt"
)

func isNotEmpty(s []rune) bool {
	return len(s) != 0
}

func backspaceCompare(s string, t string) bool {
	stack1 := make([]rune, 0)
	stack2 := make([]rune, 0)
	for _, v := range s {
		if v == '#' && isNotEmpty(stack1) {
			stack1 = stack1[:len(stack1)-1]
		} else if v != '#' {
			stack1 = append(stack1, v)

		}
	}
	for _, v := range t {
		if v == '#' && isNotEmpty(stack2) {
			stack2 = stack2[:len(stack2)-1]
		} else if v != '#' {
			stack2 = append(stack2, v)
		}
	}
	return string(stack1) == string(stack2)
}

// * Given two strings s and t, return true if they are equal when both are typed into empty text editors. '#' means a backspace character.
// * Note that after backspacing an empty text, the text will continue empty.
func main() {
	var (
		s = "ab#c"
		t = "ad#c"
	)
	// fmt.Println(backspaceCompare(s, t))
	// s, t = "ab##", "c#d#"
	// fmt.Println(backspaceCompare(s, t))
	// s, t = "a#c", "b"
	// fmt.Println(backspaceCompare(s, t))
	s, t = "y#fo##f", "y#f#o##f"
	// * s = f
	// * t = f
	fmt.Println(backspaceCompare(s, t))
}
