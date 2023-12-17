package containsduplicate217

// improved brute-force
// Time: O(n^2) and Space: O(1)

// n the worst-case scenario, where all elements of the array are distinct, the inner loop will iterate n-i-1 times for each iteration
// of the outer loop. This leads to a total of approximately n * (n-1) / 2 comparisons, which simplifies to O(n^2).

func containsDuplicateIBF(nums []int) bool {
	n := len(nums)
	// edge case 1: empty or 1 element array
	if n < 1 {
		return false
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}

	return false
}
