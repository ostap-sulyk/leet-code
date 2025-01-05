package removeduplicatesfromsortedarray2

func removeDuplicates(nums []int) int {
	k, cnt := 0, 0
	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i] != nums[i-1] {
			nums[k] = nums[i]
			cnt = 1
			k++
		} else if cnt < 2 {
			nums[k] = nums[i]
			cnt++
			k++
		}
	}
	return k
}
