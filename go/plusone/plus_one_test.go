package plusone

import (
	"leetcode/testutil"
	"testing"
)

var cases = []testutil.Case{
	{Input: []int{1, 2, 3}, Output: []int{1, 2, 4}},
	{Input: []int{9, 9}, Output: []int{1, 0, 0}},
}

func Test_plusone(t *testing.T) {
	for _, val := range cases {
		b := testutil.Equal(val.Output.([]int), plusOne(val.Input.([]int)))
		if !b {
			t.Fatalf("input: %v output: %v message: %s", val.Input, val.Output, val.Message)
		}
	}
}
