package eventworker

import (
	"fmt"

	"github.com/LukmanulHakim18/time2go/config/logger"
	"github.com/LukmanulHakim18/time2go/model"
)

func HandleProcess(event model.Event) {
	err := fmt.Errorf("error")
	if err != nil {
		logger.GetLogger().Error("error.hit")
		HandleRetry(event)
	}
}

func HandleRetry(event model.Event) error {
	return nil
}
