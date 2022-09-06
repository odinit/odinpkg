package gtype

import (
	"fmt"
	"strings"
)

type MSS = map[string]string
type MSI = map[string]int
type MSA = map[string]any
type MIS = map[int]string
type MII = map[int]int
type MIA = map[int]any
type MAS = map[any]string
type MAI = map[any]int
type MAA = map[any]any

func JoinMSS(m MSS, sep1, sep2 string) (s string) {
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

func JoinMSI(m MSI, sep1, sep2 string) (s string) {
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

func JoinMSA(m MSA, sep1, sep2 string) (s string) {
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

func JoinMIS(m MIS, sep1, sep2 string) (s string) {
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

func JoinMII(m MII, sep1, sep2 string) (s string) {
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

func JoinMIA(m MIA, sep1, sep2 string) (s string) {
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

func JoinMAS(m MAS, sep1, sep2 string) (s string) {
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

func JoinMAI(m MAI, sep1, sep2 string) (s string) {
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

func JoinMAA(m MAA, sep1, sep2 string) (s string) {
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
