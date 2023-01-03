package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Jwt struct {
		AccessSecret string
		AccessExpire int64
	}
	Mysql struct { // mysql数据库访问配置
		DataSource string
	}
	CacheRedis cache.CacheConf // redis缓存配置
}
