package main

import (
	"fmt"

	binarysearch "neetcode/binary_search"
)

func main() {
	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	res := binarysearch.SearchMatrix(matrix, 60)
	fmt.Println(res)
}
