package leetcode_383

func CaseOne(ransomNote, magazine string) bool {
	m := make(map[int32]uint32, 24)
	for _, word := range magazine {
		m[word] += 1
	}

	for _, word := range ransomNote {
		count, ok := m[word]
		if count <= 0 || !ok {
			return false
		}

		m[word] -= 1
	}

	return true

}
