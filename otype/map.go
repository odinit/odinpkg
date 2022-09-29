package otype

import (
	"fmt"
	"strings"
)

func JoinMSS(m map[string]string, sep1, sep2 string) (s string) {
	if len(m) == 0 {
		return ""
	}
	var b strings.Builder
	for k, v := range m {
		b.WriteString(fmt.Sprintf("%s%s%s%s", k, sep1, v, sep2))
	}

	s = b.String()
	return s[:len(s)-len(sep2)]
}

func JoinMSI(m map[string]int, sep1, sep2 string) (s string) {
	if len(m) == 0 {
		return ""
	}
	var b strings.Builder
	for k, v := range m {
		b.WriteString(fmt.Sprintf("%s%s%d%s", k, sep1, v, sep2))
	}

	s = b.String()
	return s[:len(s)-len(sep2)]
}

func JoinMSA(m map[string]any, sep1, sep2 string) (s string) {
	if len(m) == 0 {
		return ""
	}
	var b strings.Builder
	for k, v := range m {
		b.WriteString(fmt.Sprintf("%s%s%v%s", k, sep1, v, sep2))
	}

	s = b.String()
	return s[:len(s)-len(sep2)]
}

func JoinMIS(m map[int]string, sep1, sep2 string) (s string) {
	if len(m) == 0 {
		return ""
	}
	var b strings.Builder
	for k, v := range m {
		b.WriteString(fmt.Sprintf("%d%s%s%s", k, sep1, v, sep2))
	}

	s = b.String()
	return s[:len(s)-len(sep2)]
}

func JoinMII(m map[int]int, sep1, sep2 string) (s string) {
	if len(m) == 0 {
		return ""
	}
	var b strings.Builder
	for k, v := range m {
		b.WriteString(fmt.Sprintf("%d%s%d%s", k, sep1, v, sep2))
	}

	s = b.String()
	return s[:len(s)-len(sep2)]
}

func JoinMIA(m map[int]any, sep1, sep2 string) (s string) {
	if len(m) == 0 {
		return ""
	}
	var b strings.Builder
	for k, v := range m {
		b.WriteString(fmt.Sprintf("%d%s%v%s", k, sep1, v, sep2))
	}

	s = b.String()
	return s[:len(s)-len(sep2)]
}

func JoinMAS(m map[any]string, sep1, sep2 string) (s string) {
	if len(m) == 0 {
		return ""
	}
	var b strings.Builder
	for k, v := range m {
		b.WriteString(fmt.Sprintf("%v%s%s%s", k, sep1, v, sep2))
	}

	s = b.String()
	return s[:len(s)-len(sep2)]
}

func JoinMAI(m map[any]int, sep1, sep2 string) (s string) {
	if len(m) == 0 {
		return ""
	}
	var b strings.Builder
	for k, v := range m {
		b.WriteString(fmt.Sprintf("%v%s%d%s", k, sep1, v, sep2))
	}

	s = b.String()
	return s[:len(s)-len(sep2)]
}

func JoinMAA(m map[any]any, sep1, sep2 string) (s string) {
	if len(m) == 0 {
		return ""
	}
	var b strings.Builder
	for k, v := range m {
		b.WriteString(fmt.Sprintf("%v%s%v%s", k, sep1, v, sep2))
	}

	s = b.String()
	return s[:len(s)-len(sep2)]
}
