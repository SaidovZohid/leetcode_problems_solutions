package main

import (
	"fmt"
<<<<<<< HEAD
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
=======
	"log"
	"sort"
	"strconv"
)

func largestNumber(nums []int) string {
	str := make([]string, len(nums))
	for i, v := range nums {
		str[i] = strconv.Itoa(v)
	}
	sort.Slice(str, func(i, j int) bool {
		log.Println(i, j)
		return str[i]+str[j] > str[j]+str[i]
	})
	result := ""
	for _, s := range str {
		result += s
	}

	allZeros := true
	for _, c := range result {
		if c != '0' {
			allZeros = false
			break
		}
	}
	if allZeros {
		return "0"
	}

	return result
>>>>>>> a0431273d2e684f79aa7282ac0f35b4dd0988aca
}

func main() {
	nums := []int{10, 2}
	fmt.Println(largestNumber(nums))
	nums = []int{3, 30, 34, 5, 9}
	fmt.Println(largestNumber(nums))
}
