package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/http/pprof"

	"github.com/LukmanulHakim18/time2go/config"
	"github.com/LukmanulHakim18/time2go/config/logger"
	"github.com/LukmanulHakim18/time2go/contract"
	"github.com/LukmanulHakim18/time2go/transport"
	"github.com/LukmanulHakim18/time2go/usecase"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.elastic.co/apm/module/apmhttp/v2"
  "github.com/prometheus/client_golang/prometheus/promhttp"
)

func RunRESTServer(ctx context.Context, usecase *usecase.UseCase) *http.Server {
	restPort := fmt.Sprintf(":%s", config.GetConfig("rest_port").GetString())
	restConn, err := net.Listen("tcp", restPort)
	if err != nil {
		logger.GetLogger().FatalWithContext(ctx, fmt.Sprintf("failed to listen port: %v", err))
	}

	gwMux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &CustomMarshaler{}),
		runtime.WithErrorHandler(ErrorCustomFormat),
		runtime.WithIncomingHeaderMatcher(CustomMatcherMrg),
	)
	gwMux.HandlePath("GET", "/debug/pprof", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) { pprof.Index(w, r) })
	gwMux.HandlePath("GET", "/debug/pprof/{function}", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		switch pathParams["function"] {
		case "cmdline":
			pprof.Cmdline(w, r)
		case "profile":
			pprof.Profile(w, r)
		case "symbol":
			pprof.Symbol(w, r)
		case "trace":
			pprof.Trace(w, r)
		default:
			pprof.Index(w, r)
		}
	})
  gwMux.HandlePath("GET", "/metrics", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	  promhttp.Handler().ServeHTTP(w, r)
  })

  metricMiddleware := NewMetricMiddleware()

  restServer := &http.Server{
		Addr:    "localhost" + restPort,
		Handler: metricMiddleware.PrometheusHTTPMiddleware(apmhttp.Wrap(gwMux)),
	}

	transport := transport.NewTransport(ctx, usecase)
	contract.RegisterEventSchedulerHandlerServer(ctx, gwMux, transport)
	go restServer.Serve(restConn)
	logger.GetLogger().InfoWithContext(ctx, fmt.Sprintf("server rest listening at %v", restConn.Addr()))
	return restServer
}
