package containsduplicate217

import "slices"

// in-place sort, then linear scan.
// Time: O(n log n) and Space: O(1)

func containsDuplicateSortedApproach(nums []int) bool {
	n := len(nums)

	if n < 1 {
		return false
	}

	slices.Sort(nums) // n log n

	for i := 0; i < n-1; i++ { // n
		if nums[i] == nums[i+1] {
			return true
		}
	}

	return false
}
