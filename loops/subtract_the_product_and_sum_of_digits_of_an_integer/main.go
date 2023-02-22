package main

import (
	"fmt"
)

func subtractProductAndSum(n int) int {
	var (
		product int = 1
		sum     int
	)
	for n > 0 {
		product *= (n % 10)
		sum += (n % 10)
		n /= 10
	}

	return product - sum
}

func main() {
	n := 234
	fmt.Println(subtractProductAndSum(n))
	n = 4421
	fmt.Println(subtractProductAndSum(n))
}
