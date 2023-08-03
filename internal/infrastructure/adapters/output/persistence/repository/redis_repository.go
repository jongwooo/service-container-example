package repository

import (
	"context"
	"github.com/redis/rueidis"
)

type RedisRepository struct {
	client rueidis.Client
}

func NewRedisRepository(client rueidis.Client) *RedisRepository {
	return &RedisRepository{client: client}
}

func (r *RedisRepository) Set(c context.Context, key, value string) error {
	cmd := r.client.B().Set().Key(key).Value(value).Build()
	err := r.client.Do(c, cmd).Error()

	if err != nil {
		return err
	}
	return nil
}

func (r *RedisRepository) HSet(c context.Context, key, field, value string) error {
	cmd := r.client.B().Hset().Key(key).FieldValue().FieldValue(field, value).Build()
	err := r.client.Do(c, cmd).Error()

	if err != nil {
		return err
	}
	return nil
}

func (r *RedisRepository) Hkeys(c context.Context, key string) ([]string, error) {
	cmd := r.client.B().Hkeys().Key(key).Build()
	res, err := r.client.Do(c, cmd).AsStrSlice()

	if err != nil {
		return nil, err
	}
	return res, nil
}
