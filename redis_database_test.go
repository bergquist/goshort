package main

import (
	"testing"

	"gopkg.in/redis.v3"
)

func TestCanUseRedisIncr(t *testing.T) {
	red := redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379", Password: "", DB: 0})
	db := &RedisDatabase{
		client: red,
	}

	db.Set("dummie-key", []byte("dummie-value"))
	_, getErr := db.Get("dummie-key")
	if getErr != nil {
		t.Error("Could get value from redis")
	}

	_, incErr := db.Incr("inc-key")
	if incErr != nil {
		t.Error("Could get value from redis")
	}
}
