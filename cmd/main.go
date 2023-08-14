package main

import (
	"context"
	"fmt"
	"github.com/jongwooo/service-container-example/internal/application/ports/output"
	"github.com/jongwooo/service-container-example/internal/infrastructure/adapters/config"
	"github.com/jongwooo/service-container-example/internal/infrastructure/adapters/output/persistence/repository"
)

func main() {
	var redisClient, err = config.SetupRedis()
	var c = context.Background()
	var r output.CacheOutputPort = repository.NewRedisRepository(redisClient)

	err = r.Set(c, "octocat", "Mona the Octocat")
	if err != nil {
		panic(err)
	}

	err = r.HSet(c, "species", "octocat", "Cat and Octopus")
	if err != nil {
		panic(err)
	}

	err = r.HSet(c, "species", "dinotocat", "Dinosaur and Octopus")
	if err != nil {
		panic(err)
	}

	err = r.HSet(c, "species", "robotocat", "Cat and Robot")
	if err != nil {
		panic(err)
	}

	replies, err := r.Hkeys(c, "species")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d replies:\n", len(replies))
	for i, reply := range replies {
		fmt.Printf("  %d: %s\n", i, reply)
	}
}
