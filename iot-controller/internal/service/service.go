package service

import (
	logstash_logger "github.com/KaranJagtiani/go-logstash"
	"iotController/internal/infra"
	"iotController/internal/repository"
)

type Service struct {
	db     *repository.DataBase
	Logger *logstash_logger.Logstash
	Rabbit *infra.Rabbit
}

func NewService(db *repository.DataBase, logger *logstash_logger.Logstash, rabbit *infra.Rabbit) *Service {
	return &Service{db: db, Logger: logger, Rabbit: rabbit}
}
