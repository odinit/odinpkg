package top

import (
	"errors"
	"time"
)

var Loc = time.FixedZone("GMT", 8*3600)
var DateFormat = "2006-01-02 15:04:05"
var TimeFormat = "15:04:05"

var (
	ErrPathIsNotDir = errors.New("指定路径不是文件夹")
)
