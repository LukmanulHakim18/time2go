package eventlistener

import (
	"context"
	"sync"

	"github.com/LukmanulHakim18/time2go/pkg/eventworker"
	"github.com/LukmanulHakim18/time2go/repository"
)

type EventListener struct {
	repository *repository.Repository
	handler    *eventworker.WorkerPool
	cancel     context.CancelFunc
	wg         sync.WaitGroup
}

func NewEventListener(repo *repository.Repository, eventPool *eventworker.WorkerPool) *EventListener {
	return &EventListener{
		repository: repo,
		handler:    eventPool,
	}
}

func (el *EventListener) Start(ctx context.Context) {
	listOfListener := el.repository.Redis.GetListOfListener(ctx)
	for k, listener := range listOfListener {
		el.wg.Add(1)
		go func() {
			defer el.wg.Done()
			el.listenDB(ctx, k, listener)
		}()
	}
}

// Stop melakukan graceful shutdown
func (el *EventListener) Stop() {
	if el.cancel != nil {
		el.cancel()
	}
	el.wg.Wait()
}
