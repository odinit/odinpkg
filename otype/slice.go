package gtype

import (
	"strconv"
	"strings"
)

// SliceIntToStr 将整型切片转换成字符串切片
// param
// intSlice:待转换的int类型切片
// return
// strSlice:转换后的string类型切片
func SliceIntToStr(intSlice []int) (strSlice []string) {
	strSlice = make([]string, len(intSlice))
	for i := range intSlice {
		strSlice[i] = strconv.Itoa(intSlice[i])
	}
	return
}

// SliceStringToUpper 将字符串切片的元素都转化成大写
// param
// strSlice:待转换的字符串切片
// return
// upperSlice:转换成大写的字符串切片
func SliceStringToUpper(strSlice []string) (upperSlice []string) {
	upperSlice = make([]string, len(strSlice))
	for i := range strSlice {
		upperSlice[i] = strings.ToUpper(strSlice[i])
	}
	return
}

// SliceStringToLower 将字符串切片的元素都转化成小写
// param
// strSlice:待转换的字符串切片
// return
// lowerSlice:转换成小写的字符串切片
func SliceStringToLower(strSlice []string) (lowerSlice []string) {
	lowerSlice = make([]string, len(strSlice))
	for i := range strSlice {
		lowerSlice[i] = strings.ToUpper(strSlice[i])
	}
	return
}
