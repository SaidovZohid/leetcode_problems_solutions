package main

import "fmt"

func binarySearch(arr []int, item int) int {
	low, high := 0, len(arr) - 1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == item {
			return mid
		} else if arr[mid] < item {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return 0
}

func printNumbersWithRecursion(num int) {
	if num <= 0 {
		fmt.Println("It is over!")
	} else {
		fmt.Println(num)
		num--
		printNumbersWithRecursion(num)
	}
}

func main(){
	// arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// fmt.Println(binarySearch(arr, 2))
	printNumbersWithRecursion(10)
}