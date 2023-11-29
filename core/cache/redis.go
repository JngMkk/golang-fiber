package cache

import (
	"os"
	"sync"

	"github.com/redis/go-redis/v9"
)

var (
	redisConn      *redis.Client
	redisConnMutex sync.Mutex
)

func Connect() *redis.Client {
	redisConnMutex.Lock()
	defer redisConnMutex.Unlock()

	if redisConn == nil {
		url := os.Getenv("REDIS_URL")
		opts, err := redis.ParseURL(url)
		if err != nil {
			panic(err)
		}

		redisConn = redis.NewClient(opts)
	}

	return redisConn
}
