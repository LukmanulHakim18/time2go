package repoiface

import (
	"context"
)

type Redis interface {
	HealthCheck(ctx context.Context) error
}
