package binarysearch

type TimeMap struct {
	m map[string][]pair
}

type pair struct {
	timestamp int
	value     string
}

func Constructor() TimeMap {
	return TimeMap{m: map[string][]pair{}}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.m[key] = append(this.m[key], pair{timestamp: timestamp, value: value})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if _, found := this.m[key]; !found {
		return ""
	}

	pairs := this.m[key]
	l, r := 0, len(pairs)-1

	for l <= r {
		mid := (l + r) / 2
		if pairs[mid].timestamp <= timestamp {
			if mid == len(pairs)-1 || pairs[mid+1].timestamp > timestamp {
				return pairs[mid].value
			}
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ""
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
