package lolps

import (
	"log"

	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/os/gcache"
)

func InitRedis() {
	var redisConfig = &gredis.Config{
		Address: "127.0.0.1:6379",
	}
	redis, err := gredis.New(redisConfig)
	if err != nil {
		log.Panic("redis 初始化失败，将使用内存缓存", err)
		return
	}
	champCache.SetAdapter(gcache.NewAdapterRedis(redis))
	itemCache.SetAdapter(gcache.NewAdapterRedis(redis))
	log.Println("redis 初始化成功，将使用redis缓存")
}
