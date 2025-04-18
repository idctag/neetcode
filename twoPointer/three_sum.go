package twopointer

import (
	"slices"
)

func ThreeSum(nums []int) [][]int {
	res := [][]int{}
	slices.Sort(nums)

	for l := range nums {
		if nums[l] > 0 {
			break
		}
		if l > 0 && nums[l] == nums[l-1] {
			continue
		}
		m, r := l+1, len(nums)-1

		for m < r {
			sum := nums[l] + nums[m] + nums[r]
			if sum > 0 {
				r--
			} else if sum < 0 {
				m++
			} else {
				res = append(res, []int{nums[l], nums[m], nums[r]})
				m++
				r--
				for m < r && nums[m] == nums[m-1] {
					m++
				}
			}
		}
	}

	return res
}
