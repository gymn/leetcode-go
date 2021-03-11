package _0001_two_sum

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if i1, ok := m[target-v]; ok {
			return []int{i, i1}
		}
		m[v] = i
	}
	return nil
}
