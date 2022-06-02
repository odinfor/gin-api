package main

import (
	"flag"
	"fmt"
	"gin-api/internal/config"
	"gin-api/internal/pkg/ginzap"
	"gin-api/internal/router"
	"github.com/gin-gonic/gin"
	"os"
)

var (
	env string
)

func main() {
	flag.StringVar(&env, "-f", "etc/dev.yaml", "")

	path, _ := os.Getwd()
	logFile := fmt.Sprintf("%s/%s.log", path, config.Srv.Name)
	if err := ginzap.InitLogger(logFile); err != nil {
		panic(err)
	}

	gin.SetMode(config.Srv.Release)
	engine := gin.Default()
	engine.Use(ginzap.GinLogger(), ginzap.GinRecovery(true))
	//handler.InitServiceContext()
	router.RegisterRouter(engine) // 注册路由

	fmt.Println("server success start!")

	_ = engine.Run(fmt.Sprintf("%s:%d", config.Srv.Address, config.Srv.Port)) // 启动
}
