package medium

type arrangement struct {
	currRowSpaceLeft int
	currRowMaxHeight int
	heightTillNow    int
}

func MinHeightShelves(books [][]int, shelfWidth int) int {
	queue := []arrangement{{0, 0, 0}}
	minHeight := 0

	for _, bookDimensions := range books {
		bookWidth := bookDimensions[0]
		bookHeight := bookDimensions[1]
		nextQueue := []arrangement{}

		newRowIsMust := true

		for _, currArng := range queue {
			spaceLeft := currArng.currRowSpaceLeft
			heightTillNow:=currArng.heightTillNow
			if spaceLeft < bookWidth { //not enough space in current row
				nextQueue = append(nextQueue, arrangement{
					shelfWidth - bookWidth,
					bookHeight,
					heightTillNow + bookHeight,
				})
				continue

			} else {
				newRowIsMust = false
			}


			currRowMaxHeight := currArng.currRowMaxHeight
			if currRowMaxHeight >= bookHeight { //enough space in current row, and book is shorter wrt row
				nextQueue = append(nextQueue, arrangement{
					spaceLeft - bookWidth,
					currRowMaxHeight,
					heightTillNow,
				})
			} else { //enough space in current row but height is more than max height of current row
				nextQueue = append(nextQueue, arrangement{
					spaceLeft - bookWidth,
					bookHeight,
					heightTillNow + bookHeight - currRowMaxHeight,
				})
				if heightTillNow == minHeight {
					nextQueue = append(nextQueue, arrangement{
						shelfWidth - bookWidth,
						bookHeight,
						heightTillNow + bookHeight,
					})
				}
			}
		}

		if newRowIsMust {
			minHeight += bookHeight
			queue = []arrangement{
				{shelfWidth - bookWidth, bookHeight, minHeight},
			}
		} else {
			minHeight = 1000000
			for _, currArng := range nextQueue {
				if currArng.heightTillNow < minHeight {
					minHeight = currArng.heightTillNow
				}
			}
			queue = nextQueue
		}

	}
	return minHeight
}