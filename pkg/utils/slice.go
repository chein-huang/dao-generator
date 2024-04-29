package utils

func SliceContains[T comparable](arr []T, b T, compare func(a, b T) bool) bool {
	if len(arr) == 0 {
		return false
	}
	for _, v := range arr {
		if compare != nil {
			if !compare(v, b) {
				return false
			}
		} else {
			if v != b {
				return false
			}
		}
	}
	return true
}
