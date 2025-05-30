package repoiface

import (
	"context"
	"time"

	"github.com/LukmanulHakim18/time2go/model"
	"github.com/go-redis/redis/v8"
)

type Redis interface {
	HealthCheck(ctx context.Context) error
	SetEvent(ctx context.Context, event model.Event, indexKey, triggerKey, dataKey string, releaseEvent time.Duration) error
	DeleteEvent(ctx context.Context, dbFrom int, indexKey, dataKey string) error
	GetListOfListener(ctx context.Context) map[int]*redis.PubSub
	GetDataFromDb(ctx context.Context, dbFrom int, dataKey string) (model.Event, error)
	LockEventFromDb(ctx context.Context, dbFrom int, dataKey string) error
	UnlockEventFromDb(ctx context.Context, dbFrom int, dataKey string) error
}
