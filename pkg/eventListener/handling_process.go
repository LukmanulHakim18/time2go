package eventlistener

import (
	"context"
	"log"
	"time"

	cLog "github.com/LukmanulHakim18/core/logger"
	"github.com/LukmanulHakim18/time2go/config/logger"
	"github.com/LukmanulHakim18/time2go/model"
	"github.com/LukmanulHakim18/time2go/util"
	"github.com/go-redis/redis/v8"
)

// Worker per DB Redis
func (rl *EventListener) listenDB(ctx context.Context, db int, pubsub *redis.PubSub) {
	defer pubsub.Close()

	log.Printf("[Redis DB %d] Listening expired keys...", db)

	ch := pubsub.Channel()

	for {
		select {
		case <-ctx.Done():
			log.Printf("[Redis DB %d] Listener dihentikan", db)
			return
		case msg := <-ch:
			if msg == nil {
				continue
			}
			key := msg.Payload
			if util.CheckIsEventKey(key) {

				dataKey := util.GetDataKeyFromEventKey(key)
				event, err := rl.repository.Redis.GetDataFromDb(ctx, db, dataKey)
				if err == redis.Nil {
					logger.GetLogger().Error(" Data kosong", cLog.Field{
						Key:   "dataKey",
						Value: dataKey,
					})
					continue
				} else if err != nil {
					log.Printf("[Redis DB %d] Gagal GET key %s: %v", db, dataKey, err)
					continue
				}
				go rl.processEvent(ctx, db, event)
			}
		case <-time.After(5 * time.Second):
			// prevent goroutine leak on blocking recv
		}
	}
}

func (rl *EventListener) processEvent(ctx context.Context, dbFrom int, event model.Event) {
	logData := map[string]any{
		"method": "processEvent",
		"event":  event,
	}
	// lock process
	defer rl.repository.Redis.UnlockEventFromDb(ctx, dbFrom, event.GetLockKey())
	rl.repository.Redis.LockEventFromDb(ctx, dbFrom, event.GetDataKey())
	logger.GetLogger().Debug("process_event", cLog.Field{Key: "event", Value: event})
	// Del event
	_ = rl.repository.Redis.DeleteEvent(ctx, dbFrom, event.GetIndexKey(), event.GetDataKey())

	// execute event
	resp, err := rl.repository.HttpCaller.ExecuteEvent(ctx, event.RequestConfig)
	logData["resp"] = resp
	if err != nil {
		rl.HandlingErrorProcessEvent(ctx, event)
	}
	logger.GetLogger().Info("success", cLog.ConvertMapToFields(logData)...)

	// handling process retry
}
