package main

import (
	"fmt"
)

func peakIndexInMountainArray(arr []int) int {
	l, r := 0, len(arr)-1 // * l = 2, r = 5
	for r > l {
		mid := l + (r-l)/2 // * 1, 0
		if arr[mid] < arr[mid+1] {
			l = mid + 1
		} else {
			r = mid // * 1
		}
	}

	return l
}

func main() {
	arr := []int{0, 1, 0}
	fmt.Println(peakIndexInMountainArray(arr))
	arr = []int{0, 2, 1, 0}
	fmt.Println(peakIndexInMountainArray(arr))
	arr = []int{0, 3, 5, 10, 5, 2}
	fmt.Println(peakIndexInMountainArray(arr))
}
