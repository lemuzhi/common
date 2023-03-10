package common

import (
	"go-micro.dev/v4/config"
	"log"
)

type MysqlConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Pwd      string `json:"pwd"`
	Database string `json:"database"`
	Port     int64  `json:"port"`
}

// GetMysqlFromConsul 获取mysql的配置
func GetMysqlFromConsul(config config.Config, path ...string) *MysqlConfig {
	mysqlConfig := &MysqlConfig{}
	err := config.Get(path...).Scan(mysqlConfig)
	if err != nil {
		log.Println(err)
	}
	return mysqlConfig

}
