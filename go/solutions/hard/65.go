package hard

//When submitted: Runtime 0ms, Memory 2.44MB (Beats 100%, 35%)
func IsNumber(s string) bool {
	dotMarked := false
	eMarked := false
	numAfterE := false
	digitInNum := false

	for i, c := range s {
		switch c {
		case '+', '-':
			if i == 0 || s[i-1] == 'e' || s[i-1] == 'E' {
				continue
			} else {
				return false
			}
		case '.':
			if eMarked || dotMarked {
				return false
			} else {
				dotMarked = true
			}
		case 'e', 'E':
			if !digitInNum || s[i] >= '0' && s[i] <= '9' || eMarked {
				return false
			} else {
				eMarked = true
			}
		default:
			if c >= '0' && c <= '9' {
				if eMarked {
					if !numAfterE {
						numAfterE = true
					}
				} else if !digitInNum {
					digitInNum = true

				}
			} else {
				return false
			}
		}
	}
	if eMarked && !numAfterE {
		return false
	}
	if !digitInNum {
		return false
	}
	return true
}