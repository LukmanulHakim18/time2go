package eventlistener

import (
	"context"

	"github.com/LukmanulHakim18/time2go/model"
	"github.com/LukmanulHakim18/time2go/pkg/eventworker"
	"github.com/LukmanulHakim18/time2go/repository"
)

type EventListener struct {
	repository *repository.Repository
	sendEvent  func(event model.Event)
}

func NewEventListener(repo *repository.Repository, eventPool *eventworker.WorkerPool) *EventListener {
	return &EventListener{
		repository: repo,
		sendEvent:  eventPool.SendEvent,
	}
}




func (el *EventListener) ListenEvent(ctx context.Context){

}
