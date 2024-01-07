package twosum

// 167. Two Sum II - Input Array Is Sorted
// Approach: Two Pointer
// Time: O(n) Space: O(1)
func twoSum(numbers []int, target int) []int {
	var low, sum int

	high := len(numbers) - 1

	for low < high {
		sum = numbers[low] + numbers[high]

		switch {
		case sum < target:
			low++
		case sum > target:
			high--
		default:
			return []int{low + 1, high + 1}
		}
	}

	return []int{}
}
