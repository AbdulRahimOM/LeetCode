package medium

import "sort"

/*
Used circular linked list to solve the problem
	1. Sorted the hand
	2. Created a ring(circular linked list) with groupSize nodes
	3. While iterating through the hand, I incremented the value of the nodes in the ring
	4. I ensured that the value in each node is greater than or equal to the value in the previous node. Else, I returned false
	5. Once a loop is completed, I decreased the value of each node by the value in the root node. Also, I moved the root node to the next node
		(If the loop is having the same value in all nodes, I reduced the value of all nodes to 0 and just restarted the loop for the next group)
	6. To return true, the when the last element in hand is reached, the loop should be at its end with value of all nodes being equal

	Later:
	I avoided reseting the loop or reducing all nodes by root node value (step 5), 
		because what we just need is the loop should have same value throughout at the end to be a true case
*/
type ring struct {
	root *node
}
type node struct {
	value int
	next  *node
}

// func (r *ring) reduceByRoot() {
// 	reduction := r.root.value
// 	r.root.value -= reduction
// 	current := r.root
// 	for current.next != r.root {
// 		current.next.value -= reduction
// 		current = current.next
// 	}
// }

// func (r *ring) reduceToZero() {
// 	r.root.value = 0
// 	current := r.root
// 	for current.next != r.root {
// 		current.next.value = 0
// 		current = current.next
// 	}
// }
func IsNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}
	if groupSize == 1 {
		return true
	}
	sort.Ints(hand)
	length := len(hand)

	//create ring
	ring := &ring{
		root: &node{},
	}
	current := ring.root
	for i := 1; i < groupSize; i++ {
		current.next = &node{}
		current = current.next
	}
	current.next = ring.root

	i := 0
	for i < length {
		//initialise root
		beginNum := hand[i]
		for i < length {
			if hand[i] == beginNum {
				ring.root.value++
				i++
			} else {
				// i++
				break
			}
		}

		//initialize spin
		current = ring.root
		next := beginNum + 1
		for {
			if i == length {
				if current.next.next == ring.root {
					if current.next.value == ring.root.value {
						return true
					} else {
						return false
					}
				} else {
					return false
				}
			} else if hand[i] == next {
				current.next.value++
				i++
			} else if current.value > current.next.value {
				return false
			} else {
				current = current.next
				if current.next == ring.root {
					if current.value == ring.root.value {
						// ring.reduceToZero()
						break
					} else {
						// ring.reduceByRoot()
						ring.root = ring.root.next
						next = hand[i]
						continue
					}

				} else if hand[i] > next+1 {
					return false
				} else {
					next++
				}
			}
		}
	}
	return false
}
