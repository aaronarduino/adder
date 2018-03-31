package api

// Bin4 converts a 4 bit binary number to bools
func Bin4(a, b, c, d int) []bool {
	return []bool{int2bool(a), int2bool(b), int2bool(c), int2bool(d)}
}

func int2bool(a int) bool {
	if a == 0 {
		return false
	}
	return true
}

// Bool2int converts a bool to either a 1 or 0
func Bool2int(a bool) int {
	if a == false {
		return 0
	}
	return 1
}
