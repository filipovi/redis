package redis

import (
	"github.com/filipovi/redis/config"
	redis "gopkg.in/redis.v5"
)

// Client is the Redis Client structure
type Client struct {
	*redis.Client
}

// New returns a Redis Connection
func New(path string) (*Client, error) {
	cfg, err := config.New(path)
	if err != nil {
		return nil, err
	}
	c := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.URL,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	_, err = c.Ping().Result()
	if err != nil {
		return nil, err
	}

	return &Client{c}, nil
}
