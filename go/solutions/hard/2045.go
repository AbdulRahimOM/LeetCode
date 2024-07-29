package hard

func SecondMinimum(n int, edges [][]int, time int, change int) int {
	//setting up the graph
	graph := make([][]int, n+1)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	//calculating the shortest path................(its kind of BFS approach)

	firstOneReached := false
	visitRecord := make([]int, n+1) //marks nodes with step num of the 1st visit

	checkList := make(map[int]bool) //queue list to check in each step
	checkList[1] = true             //initial step checks on node 1 only
	step := 0
	for {
		step++
		if checkList[n] { //=>arrived at the destination
			if !firstOneReached { //=>first time reaching the destination
				firstOneReached = true
			} else { //=>second time reaching the destination
				return estimateTime(step, time, change)
			}
		} else if firstOneReached { //=>once reached the destination, but not in the further one step
			// Since the destination is reachable by a reverse+forward move after first reaching n,
			// we will reach the destination again in the next step.
			// So, we can increment the step count and end here instead of proceeding further to this next step.
			return estimateTime(step+1, time, change)
		}

		//doing the to-check list
		nextCheckList := make(map[int]bool) //for next step's check list
		for i := range checkList {
			for _, num := range graph[i] {
				if visitRecord[num] == step-1 || visitRecord[num] == 0 { 
					/*considering those nodes which are visited in the previous step or not visited yet

					Q: Why taking the nodes which we first visited in the previous step?
					Ans: Because, it could be our second minimum path (with step difference 1)

					Q: Why not considering the nodes which we visited for the first time in the even more previous steps?
					Ans: Path with step difference 2(than min step path) is surely possible,
					we can simply take the case of reverse+forward again move in last step.
					*/
					if visitRecord[num] == 0 {
						visitRecord[num] = step //marking the node with the step number of the first visit
					}

					nextCheckList[num] = true //added to checklist for the next step
				}
			}
		}
		checkList = nextCheckList //updating the checkList for the next step
	}
}

func estimateTime(steps int, time int, change int) int {
	if time%change == 0 {	//this special cases reduces the time required to calculate the time
		if time%(2*change) == 0 {
			return (steps - 1) * (time)	//no red signals in path (red signal time overlaps with travel time between nodes)
		}
		return (steps-1)*(time+change) - change	//after each travel time, there is a red signal time in every node except the last node
	}
	timeNow := 0
	for i := 1; i < steps; i++ {
		if timeNow%(2*change) >= change {
			//waiting time in red signal = change - timeNow%change
			timeNow += time+ change - timeNow%change
		}else {
			timeNow += time
		}
	}
	return timeNow
}
