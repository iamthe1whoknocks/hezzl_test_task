package redis

//redis connection
import (
	"context"
	"fmt"

	"github.com/iamthe1whoknocks/hezzl_test_task/config"
	"github.com/redis/go-redis/v9"
)

// Redis struct
type Redis struct {
	Client *redis.Client
}

func New(rc *config.Redis) (*Redis, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", rc.Host, rc.Port),
		Password: rc.Password,
		DB:       rc.DB,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, fmt.Errorf("redis - NewRedis - client.Ping: %w", err)
	}
	return &Redis{
		Client: client,
	}, nil

}
