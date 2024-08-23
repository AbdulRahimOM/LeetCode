package easy

// When submitted: Runtime: 0 ms, Memory Usage: 3.24 MB (beats 100%,35%)
func ThirdMax(nums []int) int {
	queue := []int{nums[0]}
	for _, v := range nums[1:] {
		if len(queue) >= 3 {
			if queue[2] < v {
				queue = append(queue, v)
				queue = queue[1:]
			} else if queue[1] < v && queue[2] != v {
				queue[0] = queue[1]
				queue[1] = v
			} else if queue[0] < v && queue[1] != v && queue[2] != v {
				queue[0] = v
			}
		} else if len(queue) >= 2 {
			if queue[1] < v {
				queue = append(queue, v)
			} else if queue[0] < v && queue[1] != v {
				queue = append(queue, queue[1])
				queue[1] = v
			} else if queue[1]!=v && queue[0]!=v{
				queue=append(queue, queue[1])
				queue[0],queue[1]=v,queue[0]
			}
		} else {
			if queue[0] < v {
				queue = append(queue, v)
			} else if queue[0] != v {
				queue = append(queue, queue[0])
				queue[0] = v
			}
		}
	}
	if len(queue) == 3 {
		return queue[0]
	} else {
		return queue[len(queue)-1]
	}
}