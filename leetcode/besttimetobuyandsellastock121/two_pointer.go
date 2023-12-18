package besttimetobuyandsellastock121

// Two-pointers Approach
// Time: O(n) and Space: O(1)

func maxProfit(prices []int) int {
	n := len(prices)

	if n == 1 {
		return 0
	}

	var mp, cp int // current and max profits

	var l, r int // left and right pointers

	for r < n {
		cp = prices[r] - prices[l]

		if cp <= 0 {
			l = r
		}

		if cp > mp {
			mp = cp
		}

		r++
	}

	return mp
}

// Looping Strategy ( for r < n ):
// - Iterate through the prices array with the right pointer until it reaches the end.
// - Ensures that each element in the array is processed.

// Pointer Shifting:
// - Calculate currentProfit as the difference between prices[right] and prices[left].
// - If currentProfit is non-positive, set left to right, starting a new segment.
// - Update maxProfit if currentProfit is greater than the current maximum profit.
// - Increment the right pointer to move to the next element in the array.
