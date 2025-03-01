func applyOperations(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	var j int
	for i := 0; i < n-1; i++ {
		if nums[i] == nums[i+1] {
			nums[i] = nums[i] * 2
			nums[i+1] = 0
			i++
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] != 0 {
			ans[j] = nums[i]
			j++
		}
	}
	return ans
}
