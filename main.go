package main

import (
	"dousheng-backend/Middlewares"
	"dousheng-backend/Router"
	//"github.com/vanyongqi/dousheng/service",
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.Use(Middlewares.InitLogger())
	Router.InitRouter(r)
	r.Run(":8080")
	//// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	//htttp.ListenAndServe(":8080", r) 第二种启动方式，也是r.Run()封装
}