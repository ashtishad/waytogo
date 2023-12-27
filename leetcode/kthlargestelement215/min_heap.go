package kthlargestelement215

import (
	"container/heap"

	"github.com/ashtishad/waytogo/leetcode/ds"
)

// MinHeap approach: Utilize a priority queue (min-heap) to maintain the k smallest elements.
// Initialize the heap with the first k elements of the array. For each subsequent element,
// replace the smallest element in the heap if the current element is larger, resulting in the
// min-heap storing the k smallest elements at all times. The root of the heap is the kth smallest element.
//
// Time: O(n log k) and Space: O(k)
// O(k log k) for building the initial heap + O((n - k) log k) for processing the remaining elements,
// where n is the length of the input array. In the worst case (k close to n), overall time complexity is O(n log k).

func findKthLargestA2(nums []int, k int) int {
	h := &ds.MinHeap{}
	heap.Init(h)

	for _, num := range nums {
		heap.Push(h, num)

		if h.Len() > k {
			heap.Pop(h)
		}
	}

	return (*h)[0]
}
