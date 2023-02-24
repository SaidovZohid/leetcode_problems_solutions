package main

import (
	"fmt"
	"math/rand"
	"time"
)

func largestNumber(nums []int) string {
	rand.New(rand.NewSource(time.Now().UnixNano())) // Seed the random number generator

	// Create a slice with some elements
	originalSlice := []int{3, 30, 34, 5, 9}

	// Create an empty slice to hold the new elements
	var newSlice []int

	// Generate random indexes and insert the corresponding elements from the original slice
	for i := 0; i < len(originalSlice); i++ {
		randomIndex := rand.Intn(len(newSlice) + 1)
		newSlice = append(newSlice, originalSlice[i])
		copy(newSlice[randomIndex+1:], newSlice[randomIndex:])
		newSlice[randomIndex] = originalSlice[i]
	}
	
	return ""
}

func main() {
	nums := []int{10, 2}
	fmt.Println(largestNumber(nums))
	nums = []int{3, 30, 34, 5, 9}
	fmt.Println(largestNumber(nums))
}
