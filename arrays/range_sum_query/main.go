package main

import (
	"fmt"
)

type NumArray struct {
	nums []int
}

func Constructor(nums []int) NumArray {
	temp := make([]int, len(nums))
	temp[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		temp[i] += nums[i] + temp[i - 1]
	}

	return NumArray{temp}
}

func (n *NumArray) SumRange(left, right int) int {
	if left == 0 {
		return n.nums[right]
	}

	return n.nums[right] - n.nums[left - 1]
}

func main() {
	nums := []int{-2, 0, 3, -5, 2, -1}
	obj := Constructor(nums)
	fmt.Println(obj.SumRange(0, 2))
}
