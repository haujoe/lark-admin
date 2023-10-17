package main

import (
	"carot-admin/static"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	router := gin.Default()

	static.RegisterWebStatic(router)
	router.NoRoute(func(c *gin.Context) {
		c.File("web/dist/index.html")
	})
	router.GET("/api/hello", func(c *gin.Context) { // 注册路由
		timeFormat := "2006-01-02 15:04:05"
		msg := fmt.Sprintf("Hello, current time is %s", time.Now().Format(timeFormat))
		c.JSON(200, gin.H{
			"message": msg,
		})
	})

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

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
