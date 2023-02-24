package main

import "fmt"

func containsDuplicate(nums []int) bool {
	mp := make(map[int]int)
	for _, v := range nums {
		if j, ok := mp[v]; ok && j > 0 {
			return true
		}
		mp[v]++
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(containsDuplicate(nums))
	nums = []int{1, 2, 3, 4}
	fmt.Println(containsDuplicate(nums))
	nums = []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	fmt.Println(containsDuplicate(nums))
}
