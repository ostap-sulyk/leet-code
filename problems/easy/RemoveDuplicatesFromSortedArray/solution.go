package removeduplicatesfromsortedarray

func removeDuplicates(nums []int) int {
	k := 1
	largest := nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] > largest {
			nums[k], largest = nums[i], nums[i]
			k++
		}
	}
	return k
}
