package main

import "gopkg.in/redis.v3"

type Database interface {
	Set(key string, value []byte) error
	Get(key string) (string, error)
}

type RedisDatabase struct {
	client *redis.Client
}

func (this *RedisDatabase) Get(key string) (string, error) {
	return this.client.Get(key).Result()
}

func (this *RedisDatabase) Set(key string, value []byte) error {
	_, err := this.client.Set(key, value, 0).Result()
	if err != nil {
		return err
	} else {
		return nil
	}
}
