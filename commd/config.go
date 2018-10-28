package commd

import (
	"github.com/go-ini/ini"
	"path/filepath"
	"os"
	"dep/module"
	"fmt"
)

var Config *ini.File
var Path string

func InitConfig() (Err module.Error) {
	var err error
	Path, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	Config, err = ini.Load(Path + "/../config/config.ini")
	if err != nil {
		fmt.Printf("Init Config Load failed. err:%v\n", err)
		return module.Error{ErrCode: ErrorSystem, ErrMsg: err}
	}
	fmt.Printf("init config succes.")
	return module.Error{ErrCode: SuccessCode, ErrMsg: nil}
}
