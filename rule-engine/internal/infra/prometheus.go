package infra

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	InstantRulesTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "instant_rules_total",
			Help: "Total number of instant rule alerts",
		})
	DurationRulesTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "duration_rules_total",
			Help: "Total number of errors",
		})
)
