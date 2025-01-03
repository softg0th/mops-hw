package service

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func RunMetricServer(s *Service, port string) {
	http.Handle("/metrics", promhttp.Handler())
	s.Logger.Info(map[string]interface{}{"message": "Starting Prometheus metrics server"})
	s.Logger.Error(map[string]interface{}{
		"message": http.ListenAndServe(port, nil),
		"error":   true,
	})
}
