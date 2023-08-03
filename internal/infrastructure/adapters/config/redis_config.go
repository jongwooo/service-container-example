package config

import "github.com/redis/rueidis"

func SetupRedis() rueidis.Client {
	redisClient, err := rueidis.NewClient(rueidis.ClientOption{
		InitAddress: []string{"localhost:6379"},
	})
	if err != nil {
		panic(err)
	}
	return redisClient
}
