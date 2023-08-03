package output

import "context"

type CacheOutputPort interface {
	Set(c context.Context, key, value string) error
	HSet(c context.Context, key, field, value string) error
	Hkeys(c context.Context, key string) ([]string, error)
}
