package main

import (
	"fmt"
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
}

func main() {
	nums := []int{10, 2}
	fmt.Println(largestNumber(nums))
	nums = []int{3, 30, 34, 5, 9}
	fmt.Println(largestNumber(nums))
}
