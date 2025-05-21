package main

import (
	"context"
	"time"

	"github.com/LukmanulHakim18/time2go/config"
	"github.com/LukmanulHakim18/time2go/config/logger"
	"github.com/LukmanulHakim18/time2go/config/repository"
	eventlistener "github.com/LukmanulHakim18/time2go/pkg/eventListener"
	"github.com/LukmanulHakim18/time2go/server"
	"github.com/LukmanulHakim18/time2go/usecase"
	"github.com/LukmanulHakim18/time2go/util"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	config.LoadConfigMap()
	logger.LoadLogger()
	repository.LoadRepository()
}

func main() {
	ctx := context.Background()

	usecase := usecase.NewUsecase(repository.GetRepo())
	servers := map[string]util.Operation{}

	grpcServer := server.RunGRPCServer(ctx, usecase)
	servers["grpc"] = func(ctx context.Context) error {
		grpcServer.GracefulStop()
		return nil
	}

	restServer := server.RunRESTServer(ctx, usecase)
	servers["rest"] = func(ctx context.Context) error {
		return restServer.Shutdown(ctx)
	}
	eListener := eventlistener.NewEventListener(repository.GetRepo())
	eListener.Start(ctx)
	servers["event-listener"] = func(ctx context.Context) error {
		eListener.Stop()
		return nil
	}
	eListener.Stop()
	wait := util.GracefulShutdown(ctx, 5*time.Second, servers)
	<-wait
}
