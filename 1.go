func twoSum(nums []int, target int) []int {
	diffs := make(map[int]int)
	for i, num := range nums {
		diff := target - num
		if j, ok := diffs[diff]; ok {
			return []int{i, j}
		}
		diffs[num] = i
	}
	return nil
}