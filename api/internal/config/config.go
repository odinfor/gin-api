package config

import (
	"gin-api/internal/types"
	"github.com/gin-gonic/gin"
)

/*
便于运维部署和单测镜像先将配置直接写在go文件。部署构建镜像只需要二进制文件即可。
后续考虑移入nacos
*/

var (
	Srv = &types.ServerConf{
		Address: "",
		Port:    8000,
		Release: gin.ReleaseMode,
		Name:    "demo",
	}

	JWTConf = &types.JWTConf{
		SignKey:   "this is a test sign key for jwt",
		MaxExpire: 30,
	}

	MysqlConf = &types.MysqlConf{
		Host:     "localhost",
		Port:     3306,
		Username: "root",
		Password: "root",
		Database: "demo-user",
	}
)
