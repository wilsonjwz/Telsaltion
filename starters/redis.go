package starters

import (
	"github.com/wilsonj/Telsaltion/base"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"strconv"
)

var redisClient *redis.Client

func RedisClient() *redis.Client {
	base.Check(redisClient)
	return redisClient
}

type RedisStarter struct {
	base.BaseStarter
}

func (r *RedisStarter) Init(ctx base.StarterContext) {
	logrus.Info("redis init")
}

func (r *RedisStarter) Setup(ctx base.StarterContext) {
	addr := ctx.Props().GetDefault("redis.addr", "127.0.0.1:6379")
	pwd := ctx.Props().GetDefault("redis.pwd", "")
	db := ctx.Props().GetDefault("redis.db", "0")
	redisDb, _ := strconv.Atoi(db)
	c := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd,
		DB:       redisDb,
	})
	ret, err := c.Ping().Result()
	if err != nil {
		panic(err)
	}
	logrus.Info(ret)
	redisClient = c


}

func (r *RedisStarter) Stop(ctx base.StarterContext) {
	redisClient.Close()
}
