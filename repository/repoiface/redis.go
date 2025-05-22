package repoiface

import (
	"context"
	"time"

	"github.com/LukmanulHakim18/time2go/model"
	"github.com/go-redis/redis/v8"
)

type Redis interface {
	HealthCheck(ctx context.Context) error
	SetEvent(ctx context.Context, event model.Event, indexKey, triggerKey, dataKey string, releaseEvent time.Time) error
	GetListOfListener(ctx context.Context) map[int]*redis.PubSub
	GetDataFromDb(ctx context.Context, dbFrom int, dataKey string) (model.Event, error)
	DeleteFromDb(ctx context.Context, dbFrom int, dataKey string) error
	LockEventFromDb(ctx context.Context, dbFrom int, dataKey string) error
}
