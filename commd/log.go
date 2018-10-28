package commd

import (
	"github.com/cihub/seelog"
	"dep/module"
	"fmt"
)

var Logger seelog.LoggerInterface

func InitLog() module.Error {
	Logger = seelog.Disabled
	logConfigFile := Path + Config.Section("service").Key("log_config_file").MustString(ConfigDefaultLogConfigFile)
	var err error
	Logger, err = seelog.LoggerFromConfigAsFile(logConfigFile)
	seelog.ReplaceLogger(Logger)
	if err != nil {
		fmt.Errorf("init log error:%v", err)
		return module.Error{ErrCode: ErrorSystem, ErrMsg: err}
	}
	Logger.Infof("init log success. ")
	return module.Error{ErrCode: SuccessCode, ErrMsg: nil}

}
