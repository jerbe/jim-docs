package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	_ "github.com/jerbe/jim-docs/docs"

	"github.com/gin-gonic/gin"
	"github.com/jerbe/jim-docs/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

/**
  @author : Jerbe - The porter from Earth
  @time : 2023-08-10 00:54
  @describe :
*/

func loadConfig() {
	defer func() {
		if obj := recover(); obj != nil {
			log.Println(obj)
		}
	}()

	file, err := os.Open("./swagger.json")
	if err != nil {
		log.Println(err)
		return
	}
	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Println(err)
		return
	}

	if len(bytes) == 0 {
		log.Println("文件无数据.跳过")
		return
	}

	docs.SwaggerInfo.SwaggerTemplate = string(bytes)
}

func tickLoadDocs() {
	for range time.Tick(time.Second * 10) {
		loadConfig()
	}
}

func main() {
	port := "80"
	log.Println("初始化")
	go tickLoadDocs()
	log.Println("定时加载文档已启动")
	router := gin.New()
	router.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Printf("程序已运行. 监听端口:%s, http://127.0.0.1:%s\n", port, port)
	err := router.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Println(err)
		return
	}
}
