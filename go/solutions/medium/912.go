package medium

//done by heap sort
// When submitted: Runtime: 141 ms, Memory 9.8MB (Beats 61%, 69%)
func SortArray(nums []int) []int {
	h := &maxHeap{}
	h.nums=make([]int32, len(nums))
	for i, v := range nums {
		h.nums[i] = int32(v)
	}
	h.buildHeap()

	for i := len(nums) - 1; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		h.heapifyDown(0, i-1)
	}
	for i, v := range h.nums {
		nums[i] = int(v)
	}
	return nums
}

type maxHeap struct {
	nums []int32
}

func (h *maxHeap) buildHeap() {
	lastIndex := len(h.nums) - 1
	for i := (lastIndex - 1) / 2; i >= 0; i-- {
		h.heapifyDown(i, lastIndex)
	}
}

func (h *maxHeap) heapifyDown(parent int, indexLimit int) {
	var maxIndex int
	left := 2*parent + 1
	right := 2*parent + 2

	if right <= indexLimit && h.nums[right] > h.nums[parent] {
		if h.nums[left] > h.nums[right] {
			maxIndex = left
		} else {
			maxIndex = right
		}
	} else if left <= indexLimit && h.nums[left] > h.nums[parent] {
		maxIndex = left
	} else {
		return
	}

	h.nums[parent], h.nums[maxIndex] = h.nums[maxIndex], h.nums[parent]
	h.heapifyDown(maxIndex, indexLimit)
}
