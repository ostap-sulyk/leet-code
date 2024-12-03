package twosum

func TwoSum(nums []int, target int) []int {
	myMap := make(map[int]int)

	result := []int{}
	for i, n := range nums {
		dif := target - n

		if ind, found := myMap[dif]; found {
			result = []int{ind, i}
			break
		}

		myMap[n] = i
	}
	return result
}
