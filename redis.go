package hcpairing

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	redis "github.com/go-redis/redis/v8"
)

type RedisConn interface {
	SetPlaceCache(name string, value PlaceAPIResult)
	GetPlaceCache(name string) (PlaceAPIResult, error)
}

type redisConn struct {
	connection *redis.Client
}

var (
	ctx                 = context.Background()
	CacheConn RedisConn = NewRedisConn()
)

func NewRedisConn() RedisConn {
	conn := redis.NewClient(
		&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", Config.GetRedisHost(), Config.GetRedisPort()),
			Password: Config.GetRedisPassword(),
			DB:       0,
		},
	)
	instance := redisConn{
		connection: conn,
	}
	return &instance
}

func (c *redisConn) SetPlaceCache(name string, value PlaceAPIResult) {

	tmp, _ := json.Marshal(value)
	val := string(tmp)
	fmt.Println(val)
	c.connection.Set(ctx, name, val, 24*time.Hour)
}

func (c *redisConn) GetPlaceCache(name string) (PlaceAPIResult, error) {

	val, err := c.connection.Get(ctx, name).Result()
	result := PlaceAPIResult{}
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(val), &result)
	if err != nil {
		return result, err
	}
	return result, nil
}
