package twosum

import (
	"reflect"
	"testing"
)

var testCase = []struct {
	nums    []int
	target  int
	expect  []int
	message string
}{
	{
		[]int{2, 7, 11, 15},
		9,
		[]int{0, 1},
		"官方示例",
	},
}

func Test_twoSum(t *testing.T) {
	for _, c := range testCase {
		result := twoSum(c.nums, c.target)
		if !equal(c.expect, result) {
			t.Errorf("错误用例: %s", c.message)
		}
	}
}

func Benchmark_twoSum(b *testing.B) {
	nums := []int{2, 7, 11, 15}
	target := 9
	for n := 0; n < b.N; n++ {
		b.Run("Benchmark_twoSum", func(b *testing.B) {
			result := twoSum(nums, target)
			if !equal([]int{0, 1}, result) {
				b.Errorf("运行错误")
			}
		})
	}
}

func equal(expect, actually []int) bool {
	return reflect.DeepEqual(expect, actually)
}
