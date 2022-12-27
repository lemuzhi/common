package common

import (
	"github.com/go-micro/plugins/v4/config/source/consul"
	_ "github.com/spf13/viper/remote"
	"go-micro.dev/v4/config"
	"strconv"
)

// 设置配置中心
func GetConsulConfig(host string, port int64, prefix string) (config.Config, error) {
	consulSource := consul.NewSource(
		//设置配置中心的地址
		consul.WithAddress(host+":"+strconv.FormatInt(port, 10)),
		//设置前缀，不设置将会使用默认前缀（/micro/config）
		consul.WithPrefix(prefix),
		//是否移除前缀，这里是设置为true, 表示可以不带前缀直接获取对应配置
		consul.StripPrefix(true),
	)
	//配置初始化
	conf, err := config.NewConfig()
	if err != nil {
		return conf, err
	}
	//加载配置
	err = conf.Load(consulSource)
	return conf, err
	//viper从远端consul读取配置
	//vp := viper.New()
	//err := vp.AddRemoteProvider("consul", "http://121.40.63.97:8500", "micro/config/mysql")
	//if err != nil {
	//	log.Println(err, debug.Stack())
	//}
	//vp.SetConfigType("json") // because there is no file extension in a stream of bytes, supported extensions are "json", "toml", "yaml", "yml", "properties", "props", "prop", "env", "dotenv"
	//err = vp.ReadRemoteConfig()
	//if err != nil {
	//	log.Println(err, debug.Stack())
	//}
	//fmt.Println("密码", vp.Get("pwd"))
	//return nil, err
}
