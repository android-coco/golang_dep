package httpserver

import (
	"github.com/gin-gonic/gin"
	"dep/commd"
	"dep/httpserver/api"
	"github.com/gin-contrib/cors"
	"strings"
)

func Run() {
	isDebugMode := commd.Config.Section("gin").Key("debug_mode").MustBool(true)
	if !isDebugMode {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.New(initCorsConf()))
	api.InitRoutes(router)

	err := router.Run(commd.ServerAddr)
	if err != nil {
		commd.Logger.Errorf("fail to start web service: %s", err.Error())
	}
}

func initCorsConf() cors.Config {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{
		"https://eospark.com",
		"http://localhost:10000",
	}
	config.AllowOriginFunc = func(origin string) bool {
		return strings.HasSuffix(origin, ".eospark.com") ||
			strings.HasSuffix(origin, "//eospark.com") ||
			strings.HasSuffix(origin, ".blockabc.com") ||
			strings.HasSuffix(origin, "//blockabc.com")
	}
	return config
}
