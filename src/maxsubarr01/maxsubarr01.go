package maxsubarr01

import (
	"math"
)

func MaxSubarrWithEqual01(arr []int) (result []int, found bool) {
	maxLen, maxStart, maxEnd, sumSoFar := math.MinInt32, -1, -1, 0
	seen := make(map[int]int)
	replaced := replaceZeros(arr)
	for i, val := range replaced {
		sumSoFar += val
		if endIndex, found := seen[sumSoFar]; found {
			// The subarr ending here has 0 sum
			currLen := i - endIndex + 1
			if currLen > maxLen {
				maxLen = currLen
				maxStart = endIndex + 1
				maxEnd = i
			}
		} else {
			seen[sumSoFar] = i
		}
	}

	if maxLen > math.MaxInt32 {
		return nil, false
	}

	return arr[maxStart : maxEnd+1], true
}

func replaceZeros(arr []int) []int {
	replaced := make([]int, len(arr))
	for i, val := range arr {
		if val == 0 {
			replaced[i] = -1
		} else {
			replaced[i] = val
		}
	}

	return replaced
}
