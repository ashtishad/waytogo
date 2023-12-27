package ds

// To construct a heap, we must need to implement these methods.
// type Interface interface {
// sort.Interface (len, less, swap)
// Push(x any) // add x as element Len()
// Pop() any   // remove and return element Len() - 1.

// MinHeap is a min-heap implementation using a priority queue.
type MinHeap []int

func (h *MinHeap) Len() int           { return len(*h) }
func (h *MinHeap) Less(i, j int) bool { return (*h)[i] < (*h)[j] }
func (h *MinHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *MinHeap) Push(x interface{}) {
	if val, ok := x.(int); ok {
		*h = append(*h, val)
	} else {
		panic("Push: Type assertion failed")
	}
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}
