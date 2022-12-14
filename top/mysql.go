/*
初始化MySQL连接
*/

package top

import (
	"fmt"
	"github.com/odinit/odinpkg/type"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MySQLDB *gorm.DB

func NewMySQLDB(host string, port int, user, password, dbname string, param map[string]any, opt *gorm.Config, tb ...interface{}) (db *gorm.DB, err error) {
	if opt == nil {
		opt = &gorm.Config{}
	}
	db, err = gorm.Open(
		mysql.Open(DSN(host, port, user, password, dbname, param)),
		opt,
	)
	if err != nil {
		return
	}
	if len(tb) != 0 {
		err = db.AutoMigrate(tb...)
	}
	return
}

func GlobalMySQLDB(host string, port int, user, password, dbname string, param map[string]any, opt *gorm.Config, tb ...interface{}) (err error) {
	db, err := NewMySQLDB(host, port, user, password, dbname, param, opt, tb...)
	if err != nil {
		return err
	}
	MySQLDB = db
	return
}

func DSN(host string, port int, user, password, dbname string, param map[string]any) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", user, password, host, port, dbname, _type.JoinMSA(param, "=", "&"))
}

func TestMySQLDB() (err error) {
	return GlobalMySQLDB("localhost", 3306, "root", "123456", "test", nil, nil)
}

type TempTable struct {
	Id   int    `gorm:"column:id;type:int(11);not null;primary_key" json:"-"`
	Name string `gorm:"column:name;type:varchar(255);default null;unique" json:"name"`
}
