package utils

func InSliceInt(a []int, x int) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func InSlice(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
