package service

import (
	logstash_logger "github.com/KaranJagtiani/go-logstash"
	"iotController/internal/repository"
)

type Service struct {
	db     *repository.DataBase
	Logger *logstash_logger.Logstash
}

func NewService(db *repository.DataBase, logger *logstash_logger.Logstash) *Service {
	return &Service{db: db, Logger: logger}
}
