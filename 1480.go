func runningSum(nums []int) []int {
    sums := make([]int, len(nums))
    for i, j := 0, 0; i < len(nums); i, j = i+1, j + nums[i] {
        sums[i] = nums[i] + j
    }
    return sums
}