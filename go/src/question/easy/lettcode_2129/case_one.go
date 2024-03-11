package lettcode_2129

import (
	"strings"
)

// 字符串按照空格进行拆分 words， 然后在遍历对应 words
// 默认当前字符串长度小于 2 全部转换为小写， 在全部为小写的字符时。
// 判断当前字符串长度是否大于 2， 大于情况首字母转换为大写。

func caseOne(title string) string {
	words := strings.Split(title, " ")

	for i, word := range words {
		word = strings.ToLower(word)
		if len(word) > 2 {
			word = strings.ToUpper(string(word[0])) + word[1:]
		}

		words[i] = word
	}

	return strings.Join(words, " ")
}

func caseTwo(title string) string {
	words := strings.Split(title, " ")

	for i, _ := range words {
		words[i] = strings.ToLower(words[i])
		if len(words[i]) > 2 {
			words[i] = strings.ToUpper(string(words[i][0])) + words[i][1:]
		}
	}

	return strings.Join(words, " ")
}
