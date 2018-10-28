package commd

import "dep/module"

// server
var ServerAddr string
var ReadTimeout int64

func InitEnvi() module.Error {
	ServerAddr = Config.Section("service").Key("addr").
		MustString(ConfigDefaultServerAddr)
	ReadTimeout = Config.Section("service").Key("read_timeout").
		MustInt64(ConfigDefaultReadTimeout)
	Logger.Info("init envi success")
	return module.Error{ErrCode: SuccessCode, ErrMsg: nil}
}
