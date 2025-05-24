package usecase

import (
	"context"
	"time"

	cLog "github.com/LukmanulHakim18/core/logger"
	"github.com/LukmanulHakim18/time2go/config/logger"
	"github.com/LukmanulHakim18/time2go/contract"
)

// SetEvent implements transport.UseCaseContract.
func (u *UseCase) SetEvent(ctx context.Context, request *contract.Event) (response *contract.DefaultResponse, err error) {
	var (
		logData = map[string]any{
			"method":  "SetEvent",
			"request": request,
		}
	)
	logger.GetLogger().Debug("incoming_request", cLog.ConvertMapToFields(logData)...)
	event := contract.FromProtoEvent(request)
	indexKey := event.GetIndexKey()
	triggerKey := event.GetTriggerKey()
	dataKey := event.GetDataKey()
	releaseDuration := time.Until(event.ScheduleAt)
	if err = u.Repo.Redis.SetEvent(ctx, event, indexKey, triggerKey, dataKey, releaseDuration); err != nil {
		logData["err"] = err.Error()
		logger.GetLogger().Error("error.redis.setEvent", cLog.ConvertMapToFields(logData)...)
		return nil, err
	}
	logger.GetLogger().Info("success", cLog.ConvertMapToFields(logData)...)
	return contract.GetDefaultResponse(ctx, "success create event", "berhasil menyimpan event"), nil
}
