package containsduplicate217

// brute-force
// Time: O(n^2) and Space: O(1)

func containsDuplicateBF(nums []int) bool {
	n := len(nums)
	// edge case 1: empty or 1 element array
	if n < 1 {
		return false
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}

			if nums[i] == nums[j] {
				return true
			}
		}
	}

	return false
}
