package repoiface

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type Redis interface {
	HealthCheck(ctx context.Context) error
	GetListOfListener(ctx context.Context) map[int]*redis.PubSub
}
