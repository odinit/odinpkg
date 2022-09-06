package otype

import "unicode"

// IsStrContainSpecial 判断字符串中是否包含特殊字符
func IsStrContainSpecial(s string) bool {
	if s == "" {
		return false
	}

	for _, c := range s {
		if IsCharacterSpecial(c) {
			return true
		}
	}
	return false
}

func IsCharacterSpecial(c rune) bool {
	return unicode.IsPunct(c) || unicode.IsSymbol(c) || unicode.Is(unicode.Han, c)
}

func IsStrAllLetterOrDigit(s string) bool {
	if s == "" {
		return false
	}
	for _, c := range s {
		if unicode.IsDigit(c) || unicode.IsLetter(c) {
			continue
		} else {
			return false
		}
	}
	return true
}
