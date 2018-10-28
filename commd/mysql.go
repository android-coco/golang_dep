package commd

import (
	"github.com/jinzhu/gorm"
	"fmt"
	"dep/module"
	_ "github.com/go-sql-driver/mysql"
)

var gormDefaultDB *gorm.DB

type logger struct{}

func (logger) Print(v ...interface{}) {
	fmt.Println(v)
}
func InitDb() module.Error {
	addr := Config.Section("mysql").Key("addr").MustString("")
	user := Config.Section("mysql").Key("user").MustString("")
	passwd := Config.Section("mysql").Key("passwd").MustString("")
	db := Config.Section("mysql").Key("db").MustString("")
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true",
		user, passwd, addr, db)
	var err error
	gormDefaultDB, err = gorm.Open("mysql", dataSourceName)
	if err != nil {
		// gorm 会自己 ping 一次 DB
		//Logger.Error("sql.Ping command failed. err:", err,
		//	" data_source_name: ", data_source_name)
		Logger.Errorf("sql.Open command failed. err: %v data_source_name %s ", err, dataSourceName)
		return module.Error{ErrCode: ErrorMysqlCommandPing, ErrMsg: err}
	}

	gormDefaultDB.LogMode(true).SetLogger(logger{})
	Logger.Info("init db success.")
	return module.Error{ErrCode: SuccessCode, ErrMsg: nil}

}

func GetDefaultDB() *gorm.DB {
	return gormDefaultDB
}
