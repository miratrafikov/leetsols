func removeDuplicates(nums []int) int {
	left, right := 0, 1
	for right < len(nums) {
		if nums[right] == nums[left] {
			right += 1
		} else {
			left += 1
			nums[left] = nums[right]
		}
	}
	return left + 1
}
