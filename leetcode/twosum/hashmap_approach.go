package twosum

// hash-map approach
// Time: O(n) and Space: O(n)

func twoSumHashMap(nums []int, target int) []int {
	res := make([]int, 0, 2)

	seen := make(map[int]int) // key: num val: index

	for i, num := range nums {
		comp := target - num // compliment

		if _, ok := seen[comp]; ok {
			res = append(res, seen[comp], i)
			return res
		}

		seen[num] = i
	}

	return res
}
