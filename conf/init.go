package conf

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Init 从配置文件中初始化配置信息
// confPath:配置文件路径
// conf:配置变量地址
func Init(confPath string, confValue interface{}) (err error) {
	//viper.SetConfigType("yaml") //设置文件格式,不设置则通过扩展名判断

	viper.SetConfigFile(confPath) // 设置参数文件路径
	err = viper.ReadInConfig()    // 读取配置信息
	if err != nil {
		return
	}
	err = viper.Unmarshal(confValue) // 把读取到的配置信息反序列化到 conf 变量中
	if err != nil {
		return
	}

	viper.WatchConfig()                            // 监控配置文件是否发生改动
	viper.OnConfigChange(func(in fsnotify.Event) { // 配置文件发生改动时重新加载文件
		fmt.Printf("配置文件:%s发生修改,将重新载入...\n", confPath)
		err = viper.Unmarshal(confValue)
		if err != nil {
			fmt.Printf("配置文件:%s重载失败...\n", confPath)
		}
		fmt.Printf("配置文件:%s重载完成...\n", confPath)
	})
	return
}
