package kmp

func findSubString(s string, t string) int {
	move := getMoveInfo(t)

	k := 0
	for i := 0; i < len(s); i++ {
		for i < len(s) && k < len(t) && s[i] == t[k] {
			i++
			k++
		}

		if k == len(t) {
			// 找到了子串
			return i - len(t)
		}

		if i == len(s) {
			// 找遍了s，也没找到
			return -1
		}

		// 还能继续找
		if k-1 >= 0 {
			k = move[k-1] // k-1 下标对应的子串长度，即是下一次对比的索引
			i--           // 因为后续还有 i++ 这里先回退，即保证 i 不变
		}
	}

	return -1
}

// getMoveInfo
// key: index of s
// value: length of real-sub-string, also the next check index
func getMoveInfo(s string) map[int]int {
	r := make(map[int]int)
	r[0] = 0

	p := 0

	for i := 1; i < len(s); i++ {
		for (i < len(s) && p < len(s)) && (s[i] == s[p]) {
			p++

			r[i] = p

			i++
		}

		if i == len(s) {
			// 计算完毕
			break
		}

		if p == 0 {
			r[i] = 0
		} else {
			p = r[p-1]
			i--
		}
	}

	return r
}
