package majorityelement

// First Approach with hash map
//
// func majorityElement(nums []int) int {
// 	counts := make(map[int]int)
// 	halfOfN := len(nums) / 2

//		for _, num := range nums {
//			counts[num]++
//			if counts[num] > halfOfN {
//				return num
//			}
//		}
//		return 0
//	}

func majorityElement(nums []int) int {
	candidate, count := 0, 0
	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if num == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}
