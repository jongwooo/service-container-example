package config

import "github.com/redis/rueidis"

func SetupRedis() (rueidis.Client, error) {
	redisClient, err := rueidis.NewClient(rueidis.ClientOption{
		InitAddress: []string{"localhost:6379"},
	})
	if err != nil {
		return nil, err
	}
	return redisClient, nil
}
