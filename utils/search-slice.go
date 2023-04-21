package utils

func StringSearchSlice(s []string, str string) bool {
	for _, i := range s {
		if i == str {
			return true
		}
	}
	return false
}
