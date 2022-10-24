package global

import "github.com/odinit/odinpkg/conf"

// InitConfig 从配置文件中初始化配置信息
// confPath:配置文件路径
// conf:配置变量地址
func InitConfig(confPath string, confValue interface{}) (err error) {
	return conf.Init(confPath, confValue)
}
