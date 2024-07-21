package medium

var gStart, gEnd int
var gFoundStatus int8 //0 nothing, 1 start found, 2 end found, 3 both found (S first), 4 both found (E first)
var upCount int32
var gLetters []byte

func GetDirections(root *TreeNode, startValue int, destValue int) string {
	gStart, gEnd = startValue, destValue
	gFoundStatus = 0
	upCount = 0
	gLetters = []byte{}

	traverseToFindRoute(root)

	newStrByt := []byte{}
	for range upCount {
		newStrByt = append(newStrByt, 'U')
	}
	for i := len(gLetters) - 1; i >= 0; i-- {
		newStrByt = append(newStrByt, gLetters[i])
	}

	return string(newStrByt)
}

func traverseToFindRoute(n *TreeNode) bool {
	if n == nil {
		return false
	}

	left := traverseToFindRoute(n.Left)

	switch gFoundStatus {
	case 0:
		if n.Val == gStart {
			gFoundStatus = 1
			left = true
		} else if n.Val == gEnd {
			gFoundStatus = 2
			left = true
		}
		// continue to check right
	case 1:
		if n.Val == gEnd {
			if left {
				gFoundStatus = 5
				upCount++
				return false
			} else {
				gFoundStatus = 3
				return true
			}
		}
		if left {
			upCount++
		}
		// continue to check right
	case 2:
		if n.Val == gStart {
			if left {
				gFoundStatus = 5
				gLetters = append(gLetters, 'L')
				return false
			} else {
				gFoundStatus = 4
				return true
			}
		}
		if left {
			gLetters = append(gLetters, 'L')
		}
		// else continue to check right
	case 3:
		//left is obviously true
		gLetters = append(gLetters, 'L')
		return true
	case 4:
		//left is obviously true
		upCount++
		return true
	default: //case 5:
		return false
	}

	right := traverseToFindRoute(n.Right)

	switch gFoundStatus {
	case 0:
		return false
	case 1:
		if left {
			return true
		} else if right {
			upCount++
			return true
		} else {
			return false
		}

	case 2:
		if left {
			return true
		} else if right {
			gLetters = append(gLetters, 'R')
			return true
		} else {
			return false
		}
	case 3:
		//right is obviously true
		gLetters = append(gLetters, 'R')
		if left {
			gFoundStatus = 5
			return false
		} else {
			return true
		}
	case 4:
		//right is obviously true
		upCount++
		if left {
			gFoundStatus = 5
			return false
		} else {
			return true
		}
	default: //case 5
		return false
	}
}