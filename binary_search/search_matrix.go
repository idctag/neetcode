package binarysearch

func SearchMatrix(matrix [][]int, target int) bool {
	l, r := 0, len(matrix)-1
	found := []int{}
	// find the right array
	for l <= r {
		m := l + ((r - l) / 2)
		ml := matrix[m][0]
		mr := matrix[m][len(matrix[m])-1]
		if target < ml {
			r = m - 1
			continue
		} else if target > mr {
			l = m + 1
			continue
		} else {
			found = matrix[m]
			break
		}
	}
	// check if the array was found
	if len(found) == 0 {
		return false
	}
	// find the number
	l = 0
	r = len(found) - 1
	for l <= r {
		m := l + ((r - l) / 2)
		n := found[m]
		if n == target {
			return true
		} else if n < target {
			l = m + 1
		} else if n > target {
			r = m - 1
		}
	}
	return false
}
