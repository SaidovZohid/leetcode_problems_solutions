package main

import "fmt"

func judgeCircle(moves string) bool {
	var x, y int
	for _, v := range moves {
		switch v {
		case 'U':
			x++
		case 'D':
			x--
		case 'L':
			y++
		case 'R':
			y--
		}
	}
	return x == 0 && y == 0
}

func main() {
	moves := "UD"
	fmt.Println(judgeCircle(moves))
	moves = "LL"
	fmt.Println(judgeCircle(moves))
}
