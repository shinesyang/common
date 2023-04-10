package common

import (
	"github.com/asim/go-micro/plugins/config/source/consul/v3"
	"github.com/asim/go-micro/v3/config"
	"strconv"
)

func ConsulGetConfig(hostName, prefix string, port int64) (config.Config, error) {
	// 添加配置中心
	portString := strconv.FormatInt(port, 10)

	consulSource := consul.NewSource(
		consul.WithAddress(hostName+":"+portString),
		// 设置前缀,默认前缀为: "/micro/cofnig"
		consul.WithPrefix(prefix),
		// 移除前缀,表示不带前缀也能访问
		consul.StripPrefix(true),
	)

	// 配置初始化
	newConfig, err := config.NewConfig()
	if err != nil {
		return newConfig, err
	}

	//加载配置
	err = newConfig.Load(consulSource)
	return newConfig, err
}
