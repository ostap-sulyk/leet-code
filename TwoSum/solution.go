package twosum

func TwoSum(nums []int, target int) []int {
	myMap := make(map[int]int)

	for i, num := range nums {
		dif := target - num
		difInd, ok := myMap[dif]

		if ok {
			return []int{difInd, i}
		}

		myMap[num] = i
	}

	return nil
}
