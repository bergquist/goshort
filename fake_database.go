package main

import (
	"encoding/binary"
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

func (this *inMemoryRedis) Incr(key string) (int64, error) {
	this.lck.Lock()
	defer this.lck.Unlock()

	if this.m[key] == nil {
		zero := make([]byte, 1000)
		binary.BigEndian.PutUint64(zero, 0)
		this.m[key] = []byte(zero)
	}

	var c uint64 = binary.BigEndian.Uint64(this.m[key])
	c += 1
	b := make([]byte, 1000)
	binary.BigEndian.PutUint64(b, c)
	this.m[key] = b

	return int64(c), nil
}
