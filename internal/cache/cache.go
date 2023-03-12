package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/iamthe1whoknocks/hezzl_test_task/config"
	"github.com/redis/go-redis/v9"
)

type Cache struct {
	Client *redis.Client
	Config *config.Redis
}

func New(client *redis.Client, cfg *config.Redis) *Cache {
	return &Cache{
		Client: client,
		Config: cfg,
	}
}

func (c *Cache) Set(ctx context.Context, key string, value []byte) error {
	err := c.Client.Set(ctx, key, value, time.Duration(c.Config.TTL)*time.Second).Err()
	if err != nil {
		return fmt.Errorf("cache - Set - client.Set: %w", err)
	}
	return nil
}

func (c *Cache) Get(ctx context.Context, key string) ([]byte, error) {
	b, err := c.Client.Get(ctx, key).Bytes()
	if err != nil {
		return nil, fmt.Errorf("cache - Get - client.Get: %w", err)
	}
	return b, nil
}

func (c *Cache) Invalidate(ctx context.Context, key string) error {
	err := c.Client.Del(ctx, key).Err()
	if err != nil {
		return fmt.Errorf("cache - Get - client.Del: %w", err)
	}
	return nil
}
