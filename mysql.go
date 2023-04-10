package common

import "github.com/asim/go-micro/v3/config"

type MysqlConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Pwd      string `json:"pwd"`
	Port     int32  `json:"port"`
	Database string `json:"database"`
}

// 获取consul中的mysql配置
func GetMysqlConfig(config config.Config, path ...string) (*MysqlConfig, error) {
	mysqlConfig := &MysqlConfig{}

	err := config.Get(path...).Scan(mysqlConfig)

	return mysqlConfig, err
}
