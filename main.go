package main

import (
	"log"
	"net/http"

	"github.com/SW-117/project2/bootstart"
	"github.com/SW-117/project2/controller/user"
	"github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
)

func main() {
	// 服务启动前的初始化

	if err := bootstart.Start(); err != nil {
		log.Fatalln("start other config error:" + err.Error())
		return
	}

	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})

	r.POST("/regist", user.RegisteruUser)

	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	r.Run(":8000")

	seelog.Flush()

}
