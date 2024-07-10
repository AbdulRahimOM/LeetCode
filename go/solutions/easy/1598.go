package easy

func MinOperations(logs []string) int {
	var deepness int16
	for _, log := range logs {
		switch log {
		case "./": //do nothing
		case "../":
			deepness = max(0, deepness-1)
		default:
			deepness++
		}
	}
    return int(deepness)
}