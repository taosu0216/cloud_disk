package models

import (
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"xorm.io/xorm"
)

var (
	XormEngine  *xorm.Engine
	RedisEngine *redis.Client
)

func XormInit() *xorm.Engine {
	var err error
	XormEngine, err = xorm.NewEngine("mysql", "root:root@tcp(127.0.0.1:3306)/cloud_disk")
	if err != nil {
		log.Fatalln("xorm engine init err is :", err)
	}
	return XormEngine
}
func RedisInit() *redis.Client {
	RedisEngine = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return RedisEngine
}
