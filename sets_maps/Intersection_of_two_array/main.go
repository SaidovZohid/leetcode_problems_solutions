package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	mp := make(map[int]int)
	for _, v := range nums2 {
		if _, ok := mp[v]; !ok {
			mp[v] = v
		}
	}
	arr := []int{}
	for _, v := range nums1 {
		if v, ok := mp[v]; ok {
			arr = append(arr, v)
			delete(mp, v)
		} 
	}
	return arr
}

// * Given two integer arrays nums1 and nums2, return an array of their intersection.
// * Each element in the result must be unique and you may return the result in any order.
func main() {
	nums1, nums2 := []int{1, 2, 2, 1}, []int{2, 2}
	fmt.Println(intersection(nums1, nums2))
	nums1, nums2 = []int{4, 9, 5}, []int{9, 4, 9, 8, 4}
	fmt.Println(intersection(nums1, nums2))
}
