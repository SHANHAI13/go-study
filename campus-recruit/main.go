package main

import (
	"fmt"
	"log"

	"campus-recruit/config"
	"campus-recruit/router"
)

func main() {
	// TODO: 初始化数据库
	// db := dao.InitMySQL()
	// TODO: 初始化 Redis
	// rdb := dao.InitRedis()

	r := router.SetupRouter()

	fmt.Println("========================================")
	fmt.Println("  校园招聘平台 - Campus Recruit")
	fmt.Println("  Go + Gin 后端服务")
	fmt.Println("========================================")
	log.Printf("服务启动在 http://localhost%s", config.ServerPort)
	log.Println("测试接口: GET  http://localhost:8080/ping")
	log.Println("注册接口: POST http://localhost:8080/api/v1/register")
	log.Println("登录接口: POST http://localhost:8080/api/v1/login")

	if err := r.Run(config.ServerPort); err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}
}
