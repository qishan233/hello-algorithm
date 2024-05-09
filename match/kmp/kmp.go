package kmp

func KMP(t, p string) int {
	n, m := len(t), len(p)

	k := getPrefix(p)

	q := 0
	for i := 0; i < n; i++ {
		for q > 0 && p[q] != t[i] {
			q = k[q-1]
		}

		if p[q] == t[i] {
			q++
		}

		if q == m {
			return i - m + 1
		}
	}

	return -1
}

func getPrefix(s string) map[int]int {
	r := make(map[int]int)
	m := len(s)

	r[0] = 0
	k := 0

	for q := 1; q < m; q++ {
		for k > 0 && s[q] != s[k] {
			k = r[k-1]
		}

		if s[q] == s[k] {
			k++
		}

		r[q] = k
	}

	return r
}
