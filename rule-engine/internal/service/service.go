package service

import (
	logstash_logger "github.com/KaranJagtiani/go-logstash"
	"ruleEngine/internal/infra"
)

type Service struct {
	Logger *logstash_logger.Logstash
	Rabbit *infra.Rabbit
}

func NewService(logger *logstash_logger.Logstash, rabbit *infra.Rabbit) *Service {
	return &Service{Logger: logger, Rabbit: rabbit}
}
