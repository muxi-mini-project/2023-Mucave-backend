package Midware

import (
	"Mucave/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Router /user/ [any]
// Router /post/ [any]

// @Summary  身份验证
// @Description  获取token并验证，成功则将Set UserId
// @Tags Midware
// @Accept application/json
// @Produce application/json
// @Param token header string false "token"
// @Failure 401 {object} handler.Error  "{"msg":"权限不足"}"
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
