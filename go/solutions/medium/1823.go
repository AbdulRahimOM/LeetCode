package medium

func FindTheWinner(n int, k int) int {
	removeRegister := make([]bool, n)
	removeCount := 0
	now := 0
	sum := 0
	for range n - 1 {
		remainder := k % (n - removeCount)
		if remainder == 0 {
			remainder += k
		}
		for i := 1;i <= remainder;i++{
			if removeRegister[(now+i)%n] {
				remainder++
			}
		}
		now = (now + remainder) % n
		removeCount++
		removeRegister[now] = true                                                                                                     
		sum += now
	}
	if sum == (n*(n-1))/2 {
		return n
	} else {
		return (n*(n-1))/2 - sum
	}
}