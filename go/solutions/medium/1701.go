package medium

func AverageWaitingTime(customers [][]int) float64 {
	completionTime := 0
	wait := 0
	for i := range customers {
		completionTime = max(completionTime, customers[i][0]) + customers[i][1]
		wait += completionTime - customers[i][0]
	}
	return float64(wait) / float64(len(customers))
}