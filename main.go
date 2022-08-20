package main

import (
	"log"

	"github.com/SW-117/project2/bootstart"
	"github.com/SW-117/project2/router"
	"github.com/cihub/seelog"
)

func main() {
	// 服务启动前的初始化

	if err := bootstart.Start(); err != nil {
		log.Fatalln("start other config error:" + err.Error())
		return
	}

	router.Router()

	seelog.Flush()

}
