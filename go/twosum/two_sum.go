package twosum

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if ii, ok := m[target-v]; ok {
			return []int{ii, i}
		}
		m[v] = i
	}

	return nil
}
