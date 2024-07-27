package medium

/*Honestly, 
I couldn't solve this problem on my own, all efforts was breaching the time limit.
This solution is implemented using the Floyd-Warshall algorithm, which I learned from the solutions. 
A simple algorithm that finds the shortest path between all pairs of vertices in a graph.
I am sure it would be useful in many other scenarios, even beyond graph problems, so I remember it.
I further optimized the code by using smaller data types to save memory and time.
*/
func FindTheCity(n int, edges [][]int, distanceThreshold int) int {
	dist := make([][]int16, n)
	distThreshold := int16(distanceThreshold)
	for i := range n {
		dist[i] = make([]int16, n)
		for j := range dist[i] {
			dist[i][j] = 10001
		}
		dist[i][i] = 0
	}

	for _, edge := range edges {
		dist[edge[0]][edge[1]] = int16(edge[2])
		dist[edge[1]][edge[0]] = int16(edge[2])
	}

	for middle := range n {
		for start := range n {
			for end := range n {
				
				if dist[start][middle]+dist[middle][end] < dist[start][end] {
					dist[start][end] = dist[start][middle] + dist[middle][end]
				}
			}
		}
	}

	minReachCount := int8(n)
	city := 0

	
	for start := range n {
		var reachableCities int8
		for end := range n {
			if dist[start][end] <= distThreshold {
				reachableCities++
			}
		}
		if reachableCities <= minReachCount {
			minReachCount = reachableCities
			city = start
		}
	}

	return city
}