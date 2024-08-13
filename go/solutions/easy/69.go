package easy

func MySqrt(x int) int {
	incre := uint32(1)
	var i uint32
	xx := uint32(x)
	for {
		if (i+incre)*(i+incre) < xx {
			i += incre
			incre *= 2
		} else if (i+incre)*(i+incre) == xx {
			return int(i + incre)
		} else {
			if incre <= 1 {
				break
			}
			incre /= 2
		}
	}
	return int(i)
}
