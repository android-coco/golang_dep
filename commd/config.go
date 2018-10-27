package commd

import (
	"github.com/go-ini/ini"
	"path/filepath"
	"os"
	"dep/module"
)

var Config *ini.File
var Path string

func InitConfig() (Err module.Error) {
	var err error
	Path, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	Config, err = ini.Load(Path + "/../config/config.ini")
	if err != nil {
		Logger.Error("Init Config Load failed. err:%v\n", err)
		return module.Error{ErrCode: ErrorSystem, ErrMsg: err}
	}
	return module.Error{ErrCode: SuccesCode, ErrMsg: nil}
}
