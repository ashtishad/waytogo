package twosum

// brute-force
// Time: O(n^2) and Space: O(1)

func twoSum(nums []int, target int) []int {
	var res []int

	n := len(nums)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i]+nums[j] == target {
				res = append(res, i, j)
				return res // exactly one answer exists
			}
		}
	}

	return nil
}
