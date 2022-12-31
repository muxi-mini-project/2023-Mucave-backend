package Midware

import (
	"Mucave/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// TokenMiddleWare token 获取并验证
func TokenMiddleWare(c *gin.Context) {
	UserId := service.GetToken(c)
	if UserId == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
		c.Abort()
		return
	} else {
		c.Set("UserId", UserId)
		c.Next()
		//根据 id 找到对应用户信息并返回
	}
}
