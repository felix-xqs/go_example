package conf

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
	"os"
	"time"
)

type RedisConf struct {
	Addr string `mapstructure:"addr"`
	Auth string `mapstructure:"auth"`
	Db   int    `mapstructure:"db"`
}
type YRedisMap map[string]*redis.Pool

func GetRedisMap(confMap map[string]RedisConf) YRedisMap {
	var xRedisMap = YRedisMap{}
	for db, r := range confMap {
		addr := r.Addr
		redisDB := r.Db
		redisAuth := r.Auth
		rd := &redis.Pool{
			MaxIdle:     30,
			IdleTimeout: time.Minute,
			Dial: func() (conn redis.Conn, err error) {
				conn, err1 := redis.Dial("tcp", addr,
					redis.DialReadTimeout(time.Second),
					redis.DialConnectTimeout(time.Second),
				)
				if err1 != nil {
					err = fmt.Errorf("redis.Dial err: %v, req: %v", err1, addr)
					log.Println(err)
					os.Exit(1)
					return
				}

				if auth := redisAuth; auth != "" {
					if _, authErr := conn.Do("AUTH", auth); err1 != nil {
						err = fmt.Errorf("redis AUTH err: %v, req: %v,%v", authErr, addr, auth)
						log.Println(err)
						err = conn.Close()
						if err != nil {
							log.Println(err)
						}
						os.Exit(1)
						return
					}
				}
				if db := redisDB; db > 0 {
					if _, dbError := conn.Do("SELECT", db); dbError != nil {
						err = fmt.Errorf("redis SELECT err: %v, req: %v,%v", dbError, addr, db)
						log.Println(err)
						err = conn.Close()
						if err != nil {
							log.Println(err)
						}
						os.Exit(1)
						return
					}
				}
				return
			},
		}
		xRedisMap[db] = rd
	}
	return xRedisMap
}
