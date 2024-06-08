package medium

import "fmt"

type queueList struct {
	head *queueNode
	pos2 *queueNode
	tail *queueNode
}

type queueNode struct {
	player int
	next   *queueNode
}

func NewQueueList() *queueList {
	return &queueList{
		head: nil,
		tail: nil,
	}
}

func FindWinningPlayer(skills []int, k int) int {
	n := len(skills)
	queue := NewQueueList()
	marks := make([]int, n)
	player2 := &queueNode{
		player: 1,
	}
	player1 := &queueNode{
		player: 0,
		next:   player2,
	}
	queue.head = player1
	queue.pos2 = player2
	current := queue.head.next
	for i := 2; i < n; i++ {
		current.next = &queueNode{
			player: i,
		}
		current = current.next
	}
	queue.tail = current

	// queue.displayElements()

	i := 0
	max := max(n, 100)
	for {
		if i == max {
			rank1 := skills[queue.head.player]

			persistance := false
			current := queue.pos2
			for current != nil {
				if skills[current.player] > rank1 {
					persistance = true
					break
				}
				current = current.next
			}

			if !persistance {
				return queue.head.player
			}
			i = 0
		}
		if skills[queue.head.player] > skills[queue.pos2.player] {
			marks[queue.head.player]++
			if marks[queue.head.player] == k {
				return queue.head.player
			}

			queue.tail.next = queue.pos2
			queue.tail = queue.pos2

			queue.pos2 = queue.pos2.next
			queue.head.next = queue.pos2

			queue.tail.next = nil
		} else {
			// fmt.Println("+++")
			marks[queue.pos2.player]++
			if marks[queue.pos2.player] == k {
				return queue.pos2.player
			}
			queue.tail.next = queue.head
			queue.tail = queue.head

			queue.head = queue.head.next
			queue.pos2 = queue.pos2.next

			queue.tail.next = nil
		}
		// queue.displayElements()
		i++
	}
}

func (q *queueList) displayElements() {
	current := q.head
	for current != nil {
		fmt.Print(current.player, ",")
		current = current.next
	}
	fmt.Println()
}
