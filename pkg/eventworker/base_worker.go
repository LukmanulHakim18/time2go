package eventworker

import (
	"context"
	"sync"

	cLog "github.com/LukmanulHakim18/core/logger"
	"github.com/LukmanulHakim18/time2go/config"
	"github.com/LukmanulHakim18/time2go/config/logger"
	"github.com/LukmanulHakim18/time2go/model"
	"github.com/LukmanulHakim18/time2go/repository"
)

type WorkerPool struct {
	workerCount int
	jobs        chan model.Event
	wg          sync.WaitGroup
	ctx         context.Context
	cancel      context.CancelFunc
	repo        *repository.Repository
}

func NewWorkerPool(repo *repository.Repository) *WorkerPool {
	workerCount := int(config.GetConfig("worker_count").GetInt())
	queueSize := config.GetConfig("queue_size").GetInt()

	ctx, cancel := context.WithCancel(context.Background())
	return &WorkerPool{
		workerCount: workerCount,
		jobs:        make(chan model.Event, queueSize),
		ctx:         ctx,
		cancel:      cancel,
		repo:        repo,
	}
}

func (wp *WorkerPool) Start(handler func(model.Event)) {
	for i := 0; i < wp.workerCount; i++ {
		go wp.worker(i, handler)
	}
}

func (wp *WorkerPool) SendEvent(event model.Event) {
	select {
	case <-wp.ctx.Done():
		logData := map[string]any{
			"event_id": event.ID,
			"reason":   "shutdown in progress",
		}
		logger.GetLogger().Debug("event.rejected", cLog.ConvertMapToFields(logData)...)
	default:
		wp.jobs <- event
	}
}

func (wp *WorkerPool) Shutdown() {
	logData := map[string]any{
		"message": "Shutdown initiated",
	}
	logger.GetLogger().Debug("shutdown.start", cLog.ConvertMapToFields(logData)...)

	wp.cancel()
	close(wp.jobs)
	wp.wg.Wait()

	logData = map[string]any{
		"message": "Shutdown complete",
	}
	logger.GetLogger().Debug("shutdown.done", cLog.ConvertMapToFields(logData)...)
}

func (wp *WorkerPool) worker(id int, handler func(model.Event)) {
	logData := map[string]any{
		"worker_id": id,
		"status":    "started",
	}
	logger.GetLogger().Debug("worker.status", cLog.ConvertMapToFields(logData)...)

	for {
		select {
		case <-wp.ctx.Done():
			logData := map[string]any{
				"worker_id": id,
				"status":    "stopped_receiving",
			}
			logger.GetLogger().Debug("worker.status", cLog.ConvertMapToFields(logData)...)
			return
		case job, ok := <-wp.jobs:
			if !ok {
				logData := map[string]any{
					"worker_id": id,
					"status":    "channel_closed",
				}
				logger.GetLogger().Debug("worker.status", cLog.ConvertMapToFields(logData)...)
				return
			}

			wp.wg.Add(1)
			logData := map[string]any{
				"worker_id": id,
				"event_id":  job.ID,
				"status":    "processing",
			}
			logger.GetLogger().Debug("event.processing", cLog.ConvertMapToFields(logData)...)

			go func(e model.Event) {
				defer wp.wg.Done()
				handler(e)
			}(job)
		}
	}
}
