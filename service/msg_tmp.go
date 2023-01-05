package service

import (
	"Mucave/model"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func MsgTmp(c *gin.Context) model.PrivateMsg {
	sender, _ := c.Get("UserId")
	receiver := c.Param("id")
	senderId, _ := sender.(uint)
	receiverId, _ := strconv.Atoi(receiver)
	privateMsg := model.PrivateMsg{
		SenderId:   senderId,
		ReceiverId: uint(receiverId),
		Status:     1,
		SendTime:   time.Now(),
	}
	return privateMsg
}
