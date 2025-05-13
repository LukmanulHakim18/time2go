package eventlistener

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

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
			if strings.HasPrefix(key, rl.cfg.Prefix) && strings.HasSuffix(key, rl.cfg.Suffix) {
				id := strings.TrimSuffix(strings.TrimPrefix(key, rl.cfg.Prefix), rl.cfg.Suffix)
				dataKey := fmt.Sprintf("%s%s:data", rl.cfg.Prefix, id)

				data, err := rdb.Get(ctx, dataKey).Result()
				if err == redis.Nil {
					log.Printf("[Redis DB %d] Data kosong untuk key: %s", db, dataKey)
					continue
				} else if err != nil {
					log.Printf("[Redis DB %d] Gagal GET key %s: %v", db, dataKey, err)
					continue
				}

				go rl.handler(ctx, id, data)
				_ = rdb.Del(ctx, dataKey).Err()
			}
		case <-time.After(5 * time.Second):
			// prevent goroutine leak on blocking recv
		}
	}
}
