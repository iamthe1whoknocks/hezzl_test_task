package redis

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/iamthe1whoknocks/hezzl_test_task/config"
)

// Redis struct
type Redis struct {
	Client *redis.Client
}

func New(rc *config.Redis) (*Redis, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", rc.Host, rc.Port),
		Password: rc.Password,
		DB:       0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, fmt.Errorf("redis - NewRedis - client.Ping: %w", err)
	}
	return &Redis{
		Client: client,
	}, nil

}
