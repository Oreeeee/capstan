package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type EventQueue struct {
	RedisClient redis.Client
	Name        string
	Channel     *redis.PubSub
}

func NewEventQueue(redisClient redis.Client, name string) *EventQueue {
	return &EventQueue{
		RedisClient: redisClient,
		Name:        name,
		Channel:     redisClient.Subscribe(context.Background(), name),
	}
}
