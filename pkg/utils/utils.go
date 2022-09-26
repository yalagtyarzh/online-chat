package utils

func IsBetween(check, less, more int) bool {
	return less <= check && check <= more
}

func Min(i, j int) int {
	if i > j {
		return j
	}
	return i
}
