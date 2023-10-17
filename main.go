package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()

	router.ForwardedByClientIP = true
	err := router.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		log.Fatal("设置信任的代理失败")
	}

	err = router.Run(":8080")
	if err != nil {
		log.Panicln("服务器启动失败：", err.Error())
	}
}
