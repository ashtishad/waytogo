package majorityelement169

// brute-force
// Time: O(n^2) and Space: O(1)

func majorityElementBF(nums []int) int {
	n := len(nums)

	for i := 0; i <= n/2; i++ {
		cnt := 1

		for j := i + 1; j < n; j++ {
			if nums[i] == nums[j] {
				cnt++
			}
		}

		if cnt > n/2 {
			return nums[i]
		}
	}

	return -1
}
