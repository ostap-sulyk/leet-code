package mergesortedarray

func merge(nums1 []int, m int, nums2 []int, n int) {
	m--
	n--
	for i := len(nums1) - 1; i >= 0; i-- {
		if n < 0 {
			break
		}

		if m >= 0 && nums1[m] > nums2[n] {
			nums1[i] = nums1[m]
			m--
		} else {
			nums1[i] = nums2[n]
			n--
		}

	}
}
