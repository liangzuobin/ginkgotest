package cache

import (
	redis "gopkg.in/redis.v5"
)

var (
	client *redis.Client
)

// Config redis client
func Config(addr string, poolSize int) error {
	client = redis.NewClient(&redis.Options{
		Addr:     addr,
		DB:       0,
		PoolSize: poolSize,
	})
	return nil
}
