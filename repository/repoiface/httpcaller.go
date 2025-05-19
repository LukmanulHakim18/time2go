package repoiface

import (
	"context"
	"net/http"

	"github.com/LukmanulHakim18/time2go/model"
)

type HttpCaller interface {
	HealthCheck(ctx context.Context) error
	ExecuteEvent(ctx context.Context, e model.HTTPRequestConfig) (*http.Response, error)
}
