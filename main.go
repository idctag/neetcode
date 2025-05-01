package main

import (
	"fmt"

	slidingwindow "neetcode/sliding_window"
)

func main() {
	s := "abcabcbb"
	res := slidingwindow.LongestSubstring(s)
	fmt.Println(res)
}
