package main

import "unicode/utf8"

func rotateRunes(s string, shift int) string {
	n := utf8.RuneCountInString(s)
	if n == 0 {
		return ""
	}
	runes := []rune(s)
	shift = shift % n
	//прибавляем n, потому что сдвиг вправо на k позиций
	//тоже самое, что сдвиг влево на n-k позиций
	if shift < 0 {
		shift += n
	}
	result := make([]rune, n)
	for i := 0; i < n; i++ {
		result[i] = runes[(i+shift)%n]
	}
	return string(result)
}
