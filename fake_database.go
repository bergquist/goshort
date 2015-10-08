package main

import (
	"errors"
	"sync"
)

type inMemoryRedis struct {
	m   map[string][]byte
	lck sync.RWMutex
}

func NewFakeDatabase() Database {
	return &inMemoryRedis{m: make(map[string][]byte)}
}

func (this *inMemoryRedis) Get(key string) (string, error) {
	this.lck.Lock()
	defer this.lck.Unlock()

	val, ok := this.m[key]
	if !ok {
		return "", errors.New("NotFound")
	} else {
		return string(val), nil
	}
}

func (this *inMemoryRedis) Set(key string, val []byte) error {
	this.lck.Lock()
	defer this.lck.Unlock()
	this.m[key] = val
	return nil
}
