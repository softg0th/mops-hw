package infra

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	RequestsTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "grpc_requests_total",
			Help: "Total number of requests",
		})
	RequestDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "grpc_requests_duration",
		Help: "Requests duration"}, []string{"time"})
	ErrorsTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "grpc_errors_total",
			Help: "Total number of errors",
		})
)
