package api

import (
	"github.com/gin-gonic/gin"
	"dep/httpserver/api/account"
	"dep/httpserver/api/app"
	"github.com/didip/tollbooth"
	"dep/limit"
)

func InitRoutes(router *gin.Engine) {
	limiter := tollbooth.NewLimiter(1, nil)
	router.GET("/",app.Home)
	router.GET("/account/info", limit.Handler(limiter),account.GetAccountInfo)
	router.GET("/app/info",limit.Handler(limiter), app.Info)
	router.GET("/echo", app.Echo)
}
