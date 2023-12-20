package majorityelement169

// improved brute-force(sort inputs, then binary search last occurrence - first occurrence <= n/2)
// Time: (nlogn), Space O(1)
// sorting the array (nlogn) + iterate array(n) + performing binary search(logn) = nlogn

import "slices"

func majorityElement(nums []int) int {
	n := len(nums)

	if n == 1 {
		return nums[0]
	}

	slices.Sort(nums)

	for i := 0; i < n; i++ {
		// search the value
		// if found check last - first occurrence >= n/2
		if idx, ok := slices.BinarySearch(nums, nums[i]); ok {
			cnt := 1

			for j := idx + 1; j < n; j++ {
				if nums[j] == nums[j-1] {
					cnt++
				} else {
					break
				}

				if cnt > n/2 {
					return nums[idx]
				}
			}
		}
	}

	return -1
}
