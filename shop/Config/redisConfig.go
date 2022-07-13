package Config

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

func NewPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     1,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {

			return redis.Dial("tcp", "localhost:6379")
		},
	}
}
