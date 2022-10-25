package ginutil

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/odinit/odinpkg/util"
	"strings"
)

// CheckUuid
// 校验uuid格式是否正确
func CheckUuid(u string) error {
	cLen := []int{8, 4, 4, 4, 12}
	idSplit := strings.Split(u, "-")
	if len(u) != 36 || len(idSplit) != 5 {
		return errors.New("uuid格式不正确")
	}
	for i := range idSplit {
		if len(idSplit[i]) != cLen[i] {
			return errors.New("uuid格式不正确")
		}
	}
	return nil
}

// ParseBody
// 解析request body中的信息
func ParseBody(c *gin.Context, req interface{}) (err error) {
	return util.ReaderUnmarshal(c.Request.Body, req)
}
