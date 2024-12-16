package globals

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"os"
)

var RedisSession *redis.Client

func InitRedisSession() {
	RedisSession = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:6379", os.Getenv("REDIS_HOST")),
	})
}
