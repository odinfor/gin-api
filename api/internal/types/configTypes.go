package types

//
// ServerConf
// @Description: 服务启动项配置
//
type ServerConf struct {
	Address string
	Port    int
	Release string
	Name    string
}

type JWTConf struct {
	SignKey   string
	MaxExpire int
}

type MysqlConf struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}
