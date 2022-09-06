/*
初始化MySQL连接
*/

package gmysql

import (
	"fmt"
	"github.com/odinit/odinpkg/otype"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init(host string, port int, user, password, dbname string, param otype.MSA, opts ...gorm.Option) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", user, password, host, port, dbname, otype.JoinMSA("=", "&"))

	db, err = gorm.Open(
		mysql.Open(dsn),
		opts...,
	)
	return
}

func DefaultInit() (db *gorm.DB, err error) {
	return Init("localhost", 3306, "root", "123456", "test", nil)
}

func ReplaceGlobal(db *gorm.DB) {
	DB = db
}

func GlobalInit(host string, port int, user, password, dbname string, param otype.MSA, opts ...gorm.Option) (err error) {
	DB, err = Init(host, port, user, password, dbname, param, opts...)
	return
}
