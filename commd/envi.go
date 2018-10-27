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
	return module.Error{ErrCode: SuccesCode, ErrMsg: nil}
}
