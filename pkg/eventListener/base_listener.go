package eventlistener

import (
	"context"
	"sync"

	"github.com/LukmanulHakim18/time2go/repository"
)

type EventListener struct {
	repository *repository.Repository
	cancel     context.CancelFunc
	wg         sync.WaitGroup
}

func NewEventListener(repo *repository.Repository) *EventListener {
	return &EventListener{
		repository: repo,
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
