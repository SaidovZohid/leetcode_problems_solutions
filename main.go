package main

import "fmt"

func sumOfArray(arr []int, n int) int { // * {6, 1, 2, 3, 4, 5} 6 
	if n <= 0 {
		return 0
	}

	return (sumOfArray(arr, n-1) + arr[n - 1]) 
}

func main() {
	nums := []int{6, 1, 2, 3, 4, 5}
	fmt.Println(sumOfArray(nums, len(nums))) 
}
