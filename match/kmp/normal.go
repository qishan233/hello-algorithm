package kmp

func NormalSearch(s, t string) int {
	for i := 0; i <= len(s)-len(t); i++ {
		j := 0
		for ; j < len(t); j++ {
			if s[i+j] != t[j] {
				break
			}
		}
		if j == len(t) {
			return i
		}
	}
	return -1
}
