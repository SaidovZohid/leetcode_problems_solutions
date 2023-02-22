package main

import (
	"fmt"
	"math"
)

func thirdMax(nums []int) int {
	m1, m2, m3 := math.MinInt64, math.MinInt64, math.MinInt64
	for _, v := range nums {
		if v == m1 || v == m2 || v == m3 {
			continue
		}
		if v > m1 {
			m1, m2, m3 = v, m1, m2
		} else if v > m2 {
			m2, m3 = v, m2 
		} else if v > m3 {
			m3 = v
		}
	}
	if m3 != math.MinInt64 {
		return m3
	}

	return m1
}

func main() {
	nums := []int{3, 2, 1}
	fmt.Println(thirdMax(nums))
	nums = []int{1, 2}
	fmt.Println(thirdMax(nums))
	nums = []int{2, 2, 3, 1}
	fmt.Println(thirdMax(nums))
}
