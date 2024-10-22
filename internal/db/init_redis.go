package db

import (
	"context"
	"fmt"
	"time"

	"github.com/AnnonaOrg/osenv"
	"github.com/redis/go-redis/v9"
	"github.com/umfaka/umfaka_core/internal/log"
)

var RDB *redis.Client

func RedisInit() error {
	if db, err := NewRedisClient(); err != nil {
		return err
	} else {
		RDB = db
		return nil
	}
}

func NewRedisClient() (*redis.Client, error) {
	options := redis.Options{
		Addr:       osenv.GetServerDbRedisAddress(), // Redis地址
		DB:         osenv.GetServerDbRedisDB(),      // Redis库
		PoolSize:   10,                              // Redis连接池大小
		MaxRetries: 10,                              // 最大重试次数
		// ConnMaxIdleTime: time.Second * 30,                // 空闲链接超时时间
	}
	if pw := osenv.GetServerDbRedisPassword(); len(pw) > 0 {
		options.Password = pw
	}
	db := redis.NewClient(&options)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	pong, err := db.Ping(ctx).Result()
	if err == redis.Nil {
		log.Debug("[db_redis] Nil reply returned by redis db when key does not exist.")
		return nil, fmt.Errorf("[db_redis] Nil reply returned by redis db when key does not exist")
	} else if err != nil {
		log.Errorf("[db_redis] redis connRdb err,err: %v", err)
		return nil, fmt.Errorf("redis connRdb err,err: %v", err)
	} else {
		log.Debugf("[db_redis] redis connRdb success,suc: %s", pong)
		return db, nil
	}
}
func RDBClose() error {
	return RDB.Close()
}
