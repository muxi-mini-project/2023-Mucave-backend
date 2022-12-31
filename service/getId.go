package service

import "github.com/gin-gonic/gin"

func GetId(c *gin.Context) uint {
	ID, _ := c.Get("UserId")
	id, _ := ID.(uint)
	return id
}
