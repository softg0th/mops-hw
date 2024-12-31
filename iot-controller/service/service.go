package service

import "iotController/repository"

type Service struct {
	db *repository.DataBase
}

func NewService(db *repository.DataBase) *Service {
	return &Service{db: db}
}
