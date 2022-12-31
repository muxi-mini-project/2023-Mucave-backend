package main

import (
	"Mucave/model"
	"Mucave/router"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

// @title BookSystem
// @version 1.1.0
// @description BookSystemAPI
// @termsOfService http://swagger.io/terrms/
// @contact.name  big_dust
// @contact.email 3264085417@qq.com
// @host localhost
// @BasePath
// @Schemes http
func main() {
	r = gin.Default()
	model.InitDB()
	router.Register(r)
	r.Run(":8080")
}
