package account

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAccountInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
