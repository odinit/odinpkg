/*
初始化MySQL连接
*/

package mysql

import (
	"fmt"
	"github.com/odinit/odinpkg/type"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init(host string, port int, user, password, dbname string, param map[string]any, opt *gorm.Config) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", user, password, host, port, dbname, _type.JoinMSA(param, "=", "&"))

	db, err = gorm.Open(
		mysql.Open(dsn),
		opt,
	)
	return
}

func DefaultInit() (db *gorm.DB, err error) {
	return Init("localhost", 3306, "root", "123456", "test", nil, nil)
}

func ReplaceGlobal(db *gorm.DB) {
	DB = db
}

func GlobalInit(host string, port int, user, password, dbname string, param map[string]any, opt *gorm.Config) (err error) {
	DB, err = Init(host, port, user, password, dbname, param, opt)
	return
}
