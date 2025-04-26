package main

import (
	"fmt"

	binarysearch "neetcode/binary_search"
)

func main() {
	res := binarysearch.MinEatingSpeed([]int{30, 11, 23, 4, 20}, 6)
	fmt.Println(res)
}
