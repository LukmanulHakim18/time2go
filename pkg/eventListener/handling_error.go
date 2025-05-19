package eventlistener

import (
	"context"
	"fmt"
	"time"

	"github.com/LukmanulHakim18/time2go/config"
	"github.com/LukmanulHakim18/time2go/constant"
	"github.com/LukmanulHakim18/time2go/model"
)

func (rl *EventListener) HandlingErrorProcessEvent(ctx context.Context, event model.Event) error {
	retryDelay := config.GetConfig("retry_time_delay").GetDuration()
	if event.RetryPolicy.MaxAttempts <= event.RetryPolicy.AttemptCount {
		return fmt.Errorf("error_reach_limit_retry")
	}
	switch event.RetryPolicy.Type {
	case constant.RETRY_POLICY_TYPE_EXPONENTIAL:
		retryDelay = exponentialTime(retryDelay, event.RetryPolicy.AttemptCount)
	default:

	}
	indexKey := event.GetIndexKey()
	triggerKey := event.GetTriggerKey()
	dataKey := event.GetDataKey()
	rl.repository.Redis.SetEvent(ctx, event, indexKey, triggerKey, dataKey, retryDelay)
	return nil
}

func exponentialTime(baseTime time.Duration, retryCounter int) time.Duration {
	return baseTime << retryCounter // baseTime * (2^retryCounter)
}
