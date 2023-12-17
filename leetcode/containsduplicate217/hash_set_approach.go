package containsduplicate217

// Hash-Set Approach
// Time: O(n) and Space: O(n)

func containsDuplicate(nums []int) bool {
	n := len(nums)

	if n < 1 {
		return false
	}

	// make a map Seen, iterate the slice, mark element as seen.
	// while marking the elem, if it already seen, that means it's a duplicate.

	seen := make(map[int]bool)

	for _, num := range nums {
		if seen[num] {
			return true // means, already seen
		}

		seen[num] = true
	}

	return false
}
