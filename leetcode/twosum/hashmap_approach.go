package twosum

// hash-map approach
// Time: O(n) and Space: O(n)

func twoSumHashMap(nums []int, target int) []int {
	res := make([]int, 0, 2)

	seen := make(map[int]int) // key: num val: index

	for i, num := range nums {
		compl := target - num // compliment

		_, ok := seen[compl]
		if ok {
			res = append(res, seen[compl], i)
			return res
		}

		seen[num] = i
	}

	return res
}
