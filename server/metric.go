package server

import (
	"context"
	"net/http"
	"github.com/LukmanulHakim18/time2go/config"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"google.golang.org/grpc"
)

var (
	labelNames = []string{"app_name", "pod_name", "communication_type", "method", "path", "status"}
)

var (
	// gRPC request counter
	grpcRequests = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "grpc_requests_total",
			Help: "Total number of gRPC requests",
		},
		labelNames,
	)

	// gRPC request latency
	grpcLatency = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "grpc_request_duration_seconds",
			Help:    "Latency of gRPC requests",
			Buckets: prometheus.DefBuckets,
		},
		labelNames,
	)

	// REST (gRPC-Gateway) request counter
	restRequests = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of REST API requests",
		},
		labelNames,
	)

	// REST request latency
	restLatency = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Latency of REST API requests",
			Buckets: prometheus.DefBuckets,
		},
		labelNames,
	)
)

type metricMiddleware struct {
	appName string
	podName string
}

func NewMetricMiddleware() *metricMiddleware {
	appName := config.GetConfig("app_name").GetString()
	podName := config.GetConfig("pod_name").GetString()

	return &metricMiddleware{
		appName: appName,
		podName: podName,
	}
}

func (m *metricMiddleware) ObserveGRPCRequest(label prometheus.Labels, start time.Time) {
	// Update prometheus metrics
	grpcRequests.With(label).Inc()
	grpcLatency.With(label).Observe(time.Since(start).Seconds())
}

func (m *metricMiddleware) ObserverRESTRequest(label prometheus.Labels, start time.Time) {
	// Update prometheus metrics
	restRequests.With(label).Inc()
	restLatency.With(label).Observe(time.Since(start).Seconds())
}

func (m *metricMiddleware) PrometheusUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		start := time.Now()
		resp, err := handler(ctx, req)
		status := "success"
		if err != nil {
			status = "error"
		}

		label := prometheus.Labels{
			"app_name":           m.appName,
			"pod_name":           m.podName,
			"communication_type": "grpc",
			"method":             extractMethod(info.FullMethod),
			"path":               info.FullMethod,
			"status":             status,
		}

		m.ObserveGRPCRequest(label, start)
		return resp, err
	}
}

func (m *metricMiddleware) PrometheusStreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		start := time.Now()
		err := handler(srv, ss)
		status := "success"
		if err != nil {
			status = "error"
		}

		label := prometheus.Labels{
			"app_name":           m.appName,
			"pod_name":           m.podName,
			"communication_type": "grpc",
			"method":             extractMethod(info.FullMethod),
			"path":               info.FullMethod,
			"status":             status,
		}

		m.ObserveGRPCRequest(label, start)
		return err
	}
}

func (m *metricMiddleware) PrometheusHTTPMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		ww := NewResponseWriterWrapper(w)
		next.ServeHTTP(w, r)

		status := "success"
		if ww.StatusCode < 200 || ww.StatusCode >= 300 {
			status = "error"
		}

		label := prometheus.Labels{
			"app_name":           m.appName,
			"pod_name":           m.podName,
			"communication_type": "rest",
			"method":             r.Method,
			"path":               r.URL.Path,
			"status":             status,
		}

		m.ObserverRESTRequest(label, start)
	})
}

func extractMethod(fullMethod string) string {
	parts := strings.Split(fullMethod, "/")
	if len(parts) > 2 {
		return parts[2] // The actual method name
	}
	return fullMethod
}

// ResponseWriterWrapper wraps http.ResponseWriter to capture the response status code.
type ResponseWriterWrapper struct {
	http.ResponseWriter
	StatusCode int
}

// WriteHeader captures the status code before passing it to the original ResponseWriter.
func (w *ResponseWriterWrapper) WriteHeader(statusCode int) {
	w.StatusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

// NewResponseWriterWrapper initializes ResponseWriterWrapper with default status 200.
func NewResponseWriterWrapper(w http.ResponseWriter) *ResponseWriterWrapper {
	return &ResponseWriterWrapper{ResponseWriter: w, StatusCode: http.StatusOK}
}

