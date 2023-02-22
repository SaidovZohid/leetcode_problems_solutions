package main

import "fmt"

type NumMatrix struct {
	sums [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	m, n := len(matrix), len(matrix[0]) // * 5, 5
	sums := make([][]int, m+1)
	for i := range sums {
		sums[i] = make([]int, n+1)
	} // *      1 <= 5
	for i := 1; i <= m; i++ {
		// *        1 < 5
		for j := 1; j < n; j++ {
			// * 1  1           0     0          1   0           0   1          0   0
			sums[i][j] = matrix[i-1][j-1] + sums[i][j-1] + sums[i-1][j] - sums[i-1][j-1]
		}
	}

	return NumMatrix{sums}
}

func (n *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return n.sums[row2+1][col2+1] - n.sums[row1][col2+1] - n.sums[row2+1][col1] + n.sums[row1][col1]
}

func main() {
	matrix := [][]int{{3, 0, 1, 4, 2}, {5, 6, 3, 2, 1}, {1, 2, 0, 1, 5}, {4, 1, 0, 1, 7}, {1, 0, 3, 0, 5}}
	obj := Constructor(matrix)
	fmt.Println(obj.SumRegion(1, 1, 2, 2))
}
