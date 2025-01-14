package majorityelement

func majorityElement(nums []int) int {
	counts := make(map[int]int)
	halfOfN := len(nums) / 2

	for _, num := range nums {
		counts[num]++
		if counts[num] > halfOfN {
			return num
		}
	}

	return 0
}
