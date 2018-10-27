package commd

import (
	"github.com/cihub/seelog"
	"dep/module"
)

var Logger seelog.LoggerInterface

func InitLog() module.Error {
	Logger = seelog.Disabled
	logConfigFile := Path + Config.Section("service").Key("log_config_file").MustString(ConfigDefaultLogConfigFile)
	var err error
	Logger, err = seelog.LoggerFromConfigAsFile(logConfigFile)
	seelog.ReplaceLogger(Logger)
	if err != nil {
		return module.Error{ErrCode: ErrorSystem, ErrMsg: err}
	} else {
		return module.Error{ErrCode: SuccesCode, ErrMsg: nil}
	}
}



