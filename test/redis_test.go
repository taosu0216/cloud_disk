package test

import (
	"context"
	"github.com/go-redis/redis/v8"
	"testing"
	"time"
)

var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{
	Addr:     "127.0.0.1:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

func TestSetValue(t *testing.T) {
	err := rdb.Set(ctx, "name", "taosu", time.Second*30).Err()
	if err != nil {
		t.Fatal(err)
	}
}
func TestGetValue(t *testing.T) {
	val, err := rdb.Get(ctx, "name").Result()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(val)
}
