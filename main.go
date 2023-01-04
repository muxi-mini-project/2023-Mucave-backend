package main

import (
	"Mucave/config"
	"Mucave/model"
	qiniu "Mucave/pkg"
	"Mucave/router"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

// @title Mucave
// @version 1.1.0
// @description MucaveAPI
// @termsOfService http://swagger.io/terrms/
// @contact.name  big_dust
// @contact.email 3264085417@qq.com
// @host 43.138.61.49
// @BasePath /api/v1
// @Schemes http
func main() {
	config.Init("C:/Users/onlyB/GolandProjects/a/2023-Mucave-backend/conf/config.yaml", "")
	qiniu.Load()
	model.InitDB()
	r = gin.Default()
	router.Register(r)
	r.Run(":8080")
}
