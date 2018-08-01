package redis

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/filipovi/redis/config"
	redis "gopkg.in/redis.v5"
)

// Client is the Redis Client structure
type Client struct {
	*redis.Client
}

// Cacher indicates the cache functions
type Cacher interface {
	GetHashKey(data, key string) string
	Load(key string) (string, error)
	MultiLoad(keys []string) ([]interface{}, error)
	Save(key string, data []byte) error
	FlushAll() error
}

// GetHashKey returns a key for the data
func (client Client) GetHashKey(data, key string) string {
	h := sha256.New()
	h.Write([]byte(data))
	return fmt.Sprintf(key, hex.EncodeToString(h.Sum(nil)))
}

// MultiLoad returns multiple results
func (client Client) MultiLoad(keys []string) ([]interface{}, error) {
	data, err := client.MGet(keys...).Result()
	if nil != err {
		return nil, err
	}
	return data, err
}

// FlushAll flush the redis database
func (client Client) FlushAll() error {
	return client.FlushAll()
}

// Load the data from Redis
func (client Client) Load(key string) (string, error) {
	data, err := client.Get(key).Result()
	if nil != err {
		return "", err
	}

	return data, nil
}

// Save the data in Redis
func (client Client) Save(key string, data []byte) error {
	err := client.Set(key, data, 0).Err()
	if nil != err {
		return err
	}

	return nil
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
