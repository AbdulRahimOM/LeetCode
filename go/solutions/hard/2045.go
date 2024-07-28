package hard

func SecondMinimum(n int, edges [][]int, time int, change int) int {
	graph := make([][]int, n+1)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	//logic
	checkList := make(map[int]bool)
	checkList[1] = true
	firstOneReached := false
	visitRecord := make([]int, n+1)	//marks the nodes with step number of the 1st visit
	step := 1
	for {
		if checkList[n] { //==true
			if !firstOneReached {
				firstOneReached = true
			} else {
				return estimateTime(step, time, change)
			}
		} else if firstOneReached {
			step++
			return estimateTime(step, time, change)
		}

		//queue list handling
		nextList := make(map[int]bool)
		for i := range checkList {
			for _, v := range graph[i] {
				if visitRecord[v] == step || visitRecord[v] == 0 {
					nextList[v] = true
					if visitRecord[v] == 0 {
						visitRecord[v] = step + 1
					}
				}
			}
		}
		checkList = nextList
		step++
	}
}

func estimateTime(steps int, time int, change int) int {
	if time%change==0{
		if time%(2*change)==0{
			return (steps-1)*(time)
		}
		return (steps-1)*(time+change)-change
	}
	timeNow := 0
	for i := 1; i < steps; i++ {
		if timeNow%(2*change) >= change {
			timeNow = timeNow/change*change + change
		}
		timeNow += time
	}
	return timeNow
}