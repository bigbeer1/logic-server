package redisx

import (
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/stores/cache"
)

func NewRedisManager(config cache.CacheConf) *redis.ClusterClient {
	// 解析redis-config

	var addrs []string
	var passwd string
	for _, conf := range config {
		addrs = append(addrs, conf.Host)
		passwd = conf.Pass
	}

	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    addrs,
		Password: passwd,
	})

	return rdb
}
