package session

import (
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/spf13/viper"
)

var (
	client *redis.Client
	//store  = sessions.NewCookieStore([]byte())
)

func Init() {
	client = redis.NewClient(&redis.Options{
		Addr:         viper.GetString("redis.addr"),
		Password:     "",
		DB:           viper.GetInt("redis.sessionDB"),
		PoolSize:     viper.GetInt("redis.poolSize"),
		MinIdleConns: viper.GetInt("redis.minIdleConns"),
	})
	_, err := client.Ping().Result()
	if err != nil {
		log.Printf("redis connection err: %v\n", err)
		panic(err)
	}

}

func GetClient() *redis.Client {
	return client
}

func Get(key string) (string, error) {
	value, err := client.Get(key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", fmt.Errorf("the key doesn't exist")
		}
		return "", err
	}
	return value, err
}

func Set(key string, value interface{}, expiration time.Duration) error {
	_, err := client.Set(key, value, expiration).Result()
	return err
}
