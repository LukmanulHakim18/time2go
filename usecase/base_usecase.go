package usecase

import (
	"context"

	cLog "github.com/LukmanulHakim18/core/logger"
	"github.com/LukmanulHakim18/time2go/config/logger"
	"github.com/LukmanulHakim18/time2go/contract"
	"github.com/LukmanulHakim18/time2go/repository"
)

var useCasePointer *UseCase

type UseCase struct {
	Repo *repository.Repository
}

// HealthCheck implements transport.UseCaseContract.
func (u *UseCase) HealthCheck(ctx context.Context, request *contract.EmptyRequest) (response *contract.DefaultResponse, err error) {
	var (
		logData = map[string]any{
			"method":  "HealthCheck",
			"request": request,
		}
	)
	logger.GetLogger().DebugWithContext(ctx, "incoming.request", cLog.ConvertMapToFields(logData)...)
	return contract.GetDefaultResponse(ctx, "welcome to time2go service", "selamat datang di service time2go "), nil
}

func NewUsecase(repoIn *repository.Repository) *UseCase {
	if useCasePointer == nil {
		useCasePointer = &UseCase{Repo: repoIn}
	}
	return useCasePointer
}
