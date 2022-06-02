package svc

import (
	"gin-api/internal/config"
	"gin-api/internal/model"
	pkgMysql "gin-api/internal/pkg/mysql"
	"gin-api/internal/types"
	"sync"
)

type ServiceContext struct {
	DBConf        types.MysqlConf
	DemoUserModel model.DemoUserDo
}

func NewServiceContext() *ServiceContext {
	obj := pkgMysql.MysqlObject(
		config.MysqlConf.Host,
		config.MysqlConf.Username,
		config.MysqlConf.Password,
		config.MysqlConf.Database,
		config.MysqlConf.Port,
	)
	if conn, err := obj.NewMysqlClient(); err != nil {
		panic("db connect fail")
	} else {
		once := sync.Once{}
		once.Do(func() {
			conn.AutoMigrate(&model.DemoUser{})
		})
		return &ServiceContext{
			DemoUserModel: model.NewDemoUserDo(conn),
		}
	}
}
