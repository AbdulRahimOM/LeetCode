package medium

func ValueAfterKSeconds(n int, k int) int {
	a := make([]int, n)
	for i := range a {
		a[i] = 1
	}
	for i := 0; i < k; i++ {
		for i := 1; i < n; i++ {
			a[i] += a[i-1]
            a[i]%=(1e9 + 7)
		}
	}
	return a[n-1] % (1e9 + 7)
}