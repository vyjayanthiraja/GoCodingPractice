package maxsubarr01

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_MaxSubArr01_Found(t *testing.T) {
	arr := []int{0, 0, 1, 0, 1, 0, 0, 1}
	maxLenSubarr, found := MaxSubarrWithEqual01(arr)
	assert.True(t, found, "Expected to find subarry with ")
	assert.Equal(t, []int{1, 0, 1, 0, 0, 1}, maxLenSubarr)
}
