package binarysearch

import (
	"math"
)

func MinEatingSpeed(piles []int, h int) int {
	l, r := 1, 0
	for _, v := range piles {
		r = max(r, v)
	}
	res := r
	for l <= r {
		m := l + ((r - l) / 2)
		time := 0
		for _, p := range piles {
			time += int(math.Ceil(float64(p) / float64(m)))
		}
		if time < h || time == h {
			r = m - 1
			res = min(m, res)
			continue
		} else if time > h {
			l = m + 1
			continue
		}
	}
	return res
}
