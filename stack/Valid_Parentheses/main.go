package main

import "fmt"

func isValid(s string) bool {
	stack := []rune{}
	for _, v := range s {
		if v == '{' || v == '[' || v == '(' {
			fmt.Println(v)
			stack = append(stack, v)
		} else if v == '}' && len(stack) > 0 && stack[len(stack)-1] == '{' {
			fmt.Println(v)
			stack = stack[:len(stack)-1]
		} else if v == ']' && len(stack) > 0 && stack[len(stack)-1] == '[' {
			fmt.Println(v)
			stack = stack[:len(stack)-1]
		} else if v == ')' && len(stack) > 0 && stack[len(stack)-1] == '(' {
			fmt.Println(v)
			stack = stack[:len(stack)-1]
		} else {
			fmt.Println(v)
			return false
		}
	}
	return len(stack) == 0
}

// * Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
// * An input string is valid if:
// * Open brackets must be closed by the same type of brackets.
// * Open brackets must be closed in the correct order.
// * Every close bracket has a corresponding open bracket of the same type.
func main() {
	s := "()"
	fmt.Println(isValid(s))
	s = "()[]{}"
	fmt.Println(isValid(s))
	s = "(]"
	fmt.Println(isValid(s))
}
