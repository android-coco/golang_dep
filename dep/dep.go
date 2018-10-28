package main

import (
	"dep/commd"
	"fmt"
	"dep/httpserver"
)

func main() {

	// config
	err := commd.InitConfig()
	if err.ErrCode != commd.SuccessCode {
		fmt.Errorf("init config failed. err: %v", err.ErrMsg)
		return
	}

	// log
	err = commd.InitLog()
	if err.ErrCode != commd.SuccessCode {
		fmt.Errorf("init log failed. err: %v", err.ErrMsg)
		return
	}

	// envi
	err = commd.InitEnvi()
	if err.ErrCode != commd.SuccessCode {
		commd.Logger.Error("init envi failed. err: %v", err.ErrMsg)
		return
	}

	// pid
	err = commd.InitPid()
	if err.ErrCode != commd.SuccessCode {
		commd.Logger.Error("init pid failed. err: %v", err.ErrMsg)
		return
	}

	// db
	err = commd.InitDb()
	if err.ErrCode != commd.SuccessCode {
		commd.Logger.Error("init mysql failed. err: %v", err.ErrMsg)
		return
	}
	defer commd.GetDefaultDB().Close()

	err = httpserver.Run()
	if err.ErrCode != commd.SuccessCode {
		commd.Logger.Error("init server failed. err: %v", err.ErrMsg)
		return
	}
}

