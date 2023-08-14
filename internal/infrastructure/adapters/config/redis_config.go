package config

import "github.com/redis/rueidis"

func SetupRedis() (rueidis.Client, error) {
	return rueidis.NewClient(rueidis.ClientOption{
		InitAddress: []string{"localhost:6379"},
	})
}
