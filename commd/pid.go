package commd

import (
	"os"
	"fmt"
	"dep/module"
)

func InitPid() module.Error {
	var err error
	pidFilePath := Path + Config.Section("service").Key("pid_file").MustString(ConfigDefaultPidFile)
	var f *os.File
	f, err = os.OpenFile(pidFilePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		Logger.Error("open pid file failed")
		return module.Error{ErrCode: ErrorSystem, ErrMsg: err}
	}
	defer f.Close()
	f.WriteString(fmt.Sprintf("%d", os.Getpid()))
	Logger.Info("pid file init success.")
	return module.Error{ErrCode: SuccessCode, ErrMsg: nil}
}
