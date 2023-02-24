package main

import (
	"fmt"
	"log"
)

type MinStack struct {
	s  []int
	ms []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (stack *MinStack) Push(val int) {
	if len(stack.ms) == 0 || val <= stack.ms[len(stack.ms)-1] {
		stack.ms = append(stack.ms, val)
	}
	stack.s = append(stack.s, val)
}

func (stack *MinStack) Pop() {
	if stack.GetMin() == stack.Top() {
		stack.ms = stack.ms[:len(stack.ms)-1]
	}
	stack.s = stack.s[:len(stack.s)-1]
}

func (stack *MinStack) Top() int {
	log.Println(stack.s)
	return stack.s[len(stack.s)-1]
}

func (stack *MinStack) GetMin() int {
	return stack.ms[len(stack.ms)-1]
}

/*
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
/*
* Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.
* Implement the MinStack class:
* MinStack() initializes the stack object.
* void push(int val) pushes the element val onto the stack.
* void pop() removes the element on the top of the stack.
* int top() gets the top element of the stack.
* int getMin() retrieves the minimum element in the stack.
* You must implement a solution with O(1) time complexity for each function.
 */
func main() {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-1)
	fmt.Println(obj.GetMin())
	fmt.Println(obj.Top())
	obj.Pop()
	fmt.Println(obj.GetMin())
}
