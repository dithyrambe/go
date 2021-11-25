package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"

	"lbc/cache"
)

type RedisConn struct {
	Conn *redis.Client
}

func New(port string) cache.Cache {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:" + port,
	})
	return &RedisConn{
		Conn: rdb,
	}
}

func (r *RedisConn) Set(hash string, payload []byte) error {
	err := r.Conn.Set(hash, string(payload), time.Second*3).Err()
	if err != nil {
		fmt.Printf("%v", err)
		return err
	}
	return nil
}

func (r *RedisConn) Get(hash string) ([]byte, error) {
	v, err := r.Conn.Get(hash).Result()
	if err != nil {
		fmt.Printf("%q %s", v, err)
		return nil, err
	}
	return []byte(v), nil
}
