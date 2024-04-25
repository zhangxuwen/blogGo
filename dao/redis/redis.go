package redis

import (
	"blogWeb/settings"

	"github.com/go-redis/redis"
)

// 声明一个全局的rdb变量
var rdb *redis.Client

// 初始化连接
func Init(cfg settings.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     cfg.Host,     // viper.GetString("redis.host"),
		Password: cfg.Password, // viper.GetString("redis.password"),
		DB:       cfg.DB,       // viper.GetInt("redis.db"),
		PoolSize: cfg.PoolSize, // viper.GetInt("redis.pool_size"),
	})

	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return
}

func Close() {
	_ = rdb.Close()
}
