package majorityelement169

// Boyer-Moore Majority Voting Algorithm
// Time: O(n) and Space: O(1)
//
// For each element in the sequence:
// - If the counter is zero, set the current candidate to the element, and initialize the counter to 1.
// - If the counter is non-zero, increment it if the current candidate matches the element; otherwise, decrement the counter.

func majorityElement(nums []int) int {
	var candidate, count int

	for _, num := range nums {
		if count == 0 {
			candidate, count = num, 1
		} else {
			if candidate == num {
				count++
			} else {
				count--
			}
		}
	}

	return candidate
}
