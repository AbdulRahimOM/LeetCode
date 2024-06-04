package medium

type calendar struct {
	head *meetChain
}

type meetChain struct {
	start int
	end   int
	next  *meetChain
}

// Regular Expression Matching
func CountDays(days int, meetings [][]int) int {
	calendar := &calendar{
		head: &meetChain{
			start: meetings[0][0],
			end:   meetings[0][1],
		},
	}
	for _, v := range meetings {
		calendar.add(v[0], v[1])
	}
	// calendar.display()	//for seeing
	return calendar.getGaps(days)
}
func (c *calendar) getGaps(days int) int {
	current := c.head
	for current != nil {
		days -= current.end - current.start + 1
		current = current.next
	}
	return days

}
func (c *calendar) add(start, end int) {
	if end < c.head.start {
		c.head = &meetChain{
			start: start,
			end:   end,
			next:  c.head,
		}
		return
	} else if end <= c.head.end {
		if start < c.head.start {
			c.head.start = start
		}
		return
	} else if start <= c.head.end {
		if start <= c.head.start && end >= c.head.end {
			c.head.start = start
			c.head.end = end
		} else if end > c.head.end {
			c.head.end = end
		}
		current := c.head.next
		for current != nil {
			if current.start <= end {
				if current.end >= end {
					c.head.end = current.end
				}
				c.head.next = current.next
			} else {
				return
			}
			current = current.next
		}
		return
	} else if c.head.next != nil && end < c.head.next.start {
		c.head.next = &meetChain{
			start: start,
			end:   end,
			next:  c.head.next,
		}
		return
	} else if c.head.next == nil {
		c.head.next = &meetChain{
			start: start,
			end:   end,
		}
		return
	}

	current := c.head.next
	previous := c.head
	for current != nil {
		if current.start > start {
			if current.start > end {
				previous.next = &meetChain{
					start: start,
					end:   end,
					next:  current,
				}
			} else {
				previous.next.start = start

				if current.end < end {
					previous.next.end = end
					current = current.next
					for current != nil {
						if current.start <= end {
							previous.next.end = current.end
							previous.next.next = current.next
						} else {
							return
						}
						current = current.next
					}
				}
			}
			return

		} else if current.end < end {
			if current.end >= start {
				previous.next.end = end
				current = current.next
				for current != nil {
					if current.start <= end {
						if current.end > previous.next.end {
							previous.next.end = current.end
						}
						previous.next.next = current.next
					} else {
						return
					}
					current = current.next
				}
				return
			}
		} else {
			if current.end >= end && current.start <= start {
				return
			}
		}
		current = current.next
		previous = previous.next
	}
	previous.next = &meetChain{
		start: start,
		end:   end,
	}
}

//for seeing
// func (c *calendar) display() {
// 	current := c.head
// 	for current != nil {
// 		fmt.Print("{", current.start, "-", current.end, "},")
// 		current = current.next
// 	}
// 	fmt.Println("")
// }
