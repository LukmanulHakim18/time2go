package eventlistener

import (
	"context"
	"log"
	"time"

	cLog "github.com/LukmanulHakim18/core/logger"
	"github.com/LukmanulHakim18/time2go/config/logger"
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
				// locking
				 = rl.repository.Redis.LockEventFromDb(ctx, db, dataKey)
				go rl.handler.SendEvent(event)
				_ = rl.repository.Redis.DeleteFromDb(ctx, db, dataKey)
			}
		case <-time.After(5 * time.Second):
			// prevent goroutine leak on blocking recv
		}
	}
}
