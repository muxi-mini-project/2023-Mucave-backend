package handler

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
type Error struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

func SendResponse(c *gin.Context, message string, data interface{}) {
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  message,
		"data": data,
	})
}
func SendError(c *gin.Context, code int, error string) {
	c.JSON(code, gin.H{
		"code":  code,
		"error": error,
	})
}

//func GetFiles(c *gin.Context) {
//	path := c.Query("path")
//	err := service.LoadFile(c, path)
//	if err != nil {
//		SendError(c, 410, "未获取到指定的文件.")
//		return
//	}
//	SendResponse(c, "获取到指定的文件", nil)
//}
