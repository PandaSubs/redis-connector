package main

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Cache struct {
	rdb        *redis.Client
	expiration uint32
}

type Configurations struct {
	Addr       string
	Expiration uint32
}

func Init(config *Configurations) *Cache {
	var rdb *redis.Client = redis.NewClient(&redis.Options{
		Addr: config.Addr,
	})
	return &Cache{rdb: rdb, expiration: config.Expiration}
}

func (c *Cache) Ping(ctx context.Context) error {
	if err := c.rdb.Ping(ctx).Err(); err != nil {
		return err
	}
	return nil
}

func (c *Cache) Get(ctx context.Context, key string) (string, error) {
	val, err := c.rdb.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func (c *Cache) Set(ctx context.Context, key string, value interface{}) error {
	if err := c.rdb.Set(ctx, key, value, time.Duration(c.expiration)*time.Second).Err(); err != nil {
		return err
	}
	return nil
}
