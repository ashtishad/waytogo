package majorityelement169

// Hashing-approach with tracking frequency
// Time O(n) and Space O(n)

func majorityElementHashing(nums []int) int {
	seen := make(map[int]int) // k: num val, v: frequency

	n := len(nums)

	if n == 1 {
		return nums[0]
	}

	for i := 0; i < n; i++ {
		seen[nums[i]]++

		if seen[nums[i]] > n/2 {
			return nums[i]
		}
	}

	return -1
}
