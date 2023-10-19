package main

import (
	"carot-admin/static"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	"log"
	"time"

	_ "carot-admin/docs"
)

//go:generate swag init main.go -o=./docs --parseDependency --parseDepth=1

//	@title			carot-admin 后端 API 文档
//	@version		1.0.0
//	@description	carot-admin 是基于 gin、JWT 和 RBAC 的 go 后台管理系统
//	@termsOfService	https://github.com/syakod

//	@license.name	MIT
//	@license.url	https://github.com/syakod/carot-admin/blob/main/LICENSE

//	@contact.name	syakod
//	@contact.url	https://github.com/carot-admin
//	@contact.email	hi@syakod.com

//	@host		localhost:8080
//	@BasePath	/api/v1

// @securityDefinitions.apikey	Bearer
// @in							header
// @name						Authorization

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	router := gin.Default()

	static.RegisterWebStatic(router)
	router.NoRoute(func(c *gin.Context) {
		c.File("web/dist/index.html")
	})

	router.GET("/api/v1/hello", test)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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

// @Router /hello [get]
func test(c *gin.Context) { // 注册路由
	timeFormat := "2023-10-01 00:00:00"
	msg := fmt.Sprintf("Hello, current time is %s", time.Now().Format(timeFormat))
	c.JSON(200, gin.H{
		"message": msg,
	})
}
