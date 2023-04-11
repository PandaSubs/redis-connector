package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Connection struct {
	rdb        *redis.Client
	expiration uint32
}

type Configurations struct {
	Addr       string
	Expiration uint32
}

func Init(config *Configurations) *Connection {
	var rdb *redis.Client = redis.NewClient(&redis.Options{
		Addr: config.Addr,
	})
	return &Connection{rdb: rdb, expiration: config.Expiration}
}

func (c *Connection) Ping(ctx context.Context) error {
	if err := c.rdb.Ping(ctx).Err(); err != nil {
		return err
	}
	return nil
}

func (c *Connection) Get(ctx context.Context, key string) (string, error) {
	val, err := c.rdb.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func (c *Connection) Set(ctx context.Context, key string, value interface{}) error {
	if err := c.rdb.Set(ctx, key, value, time.Duration(c.expiration)*time.Second).Err(); err != nil {
		return err
	}
	return nil
}
