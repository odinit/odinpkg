package gutil

import "math"

// SlicePage 分页
func SlicePage(pageNum, pageSize, total int) (sliceStart int) {
	if pageNum <= 0 {
		pageNum = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	pageCount := int(math.Ceil(float64(total) / float64(pageSize))) // 数据可以分为几页
	if pageNum > pageCount {                                        // 最后一页或超过最后一页,则返回最后一页
		pageNum = pageCount
	}
	sliceStart = (pageNum - 1) * pageSize

	return sliceStart
}
