package limit

import (
	"github.com/didip/tollbooth/limiter"
	"github.com/gin-gonic/gin"
	"github.com/didip/tollbooth"
	"net/http"
	"dep/module"
)

func Handler(lmt *limiter.Limiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		httpError := tollbooth.LimitByRequest(lmt, c.Writer, c.Request)
		if httpError != nil {
			//c.Data(httpError.StatusCode, lmt.GetMessageContentType(), []byte(httpError.Message))
			//c.Abort()
			c.AbortWithStatusJSON(http.StatusTooManyRequests, module.ApiResp{
				ErrNo:  http.StatusTooManyRequests,
				ErrMsg: http.StatusText(http.StatusTooManyRequests),
			})
		} else {
			c.Next()
		}
	}
}
