package main

import (
	"dep/commd"
	"fmt"
	"dep/httpserver"
)

func main() {

	err := commd.InitConfig()
	if err.ErrCode != commd.SuccesCode {
		commd.Logger.Error("init config failed. err: %v", err.ErrMsg)
	}
	err = commd.InitEnvi()
	if err.ErrCode != commd.SuccesCode {
		commd.Logger.Error("init envi failed. err: %v", err.ErrMsg)
	}

	err = commd.InitLog()
	if err.ErrCode != commd.SuccesCode {
		fmt.Errorf("init log failed. err: %v", err.ErrMsg)
		return
	}
	err = commd.InitPid()
	if err.ErrCode != commd.SuccesCode {
		commd.Logger.Error("init pid failed. err: %v", err.ErrMsg)
		return
	}
	err = commd.InitDb()
	if err.ErrCode != commd.SuccesCode {
		commd.Logger.Error("init mysql failed. err: %v", err.ErrMsg)
		return
	}
	defer commd.GetDefaultDB().Close()
	httpserver.Run()
}

