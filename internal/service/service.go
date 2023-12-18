package service

import (
	"feature/internal/repository"
	"feature/pkg/helper/sid"
	"feature/pkg/jwt"
	"feature/pkg/log"
)

type Service struct {
	logger *log.Logger
	sid    *sid.Sid
	jwt    *jwt.JWT
	tm     repository.Transaction
}

func NewService(tm repository.Transaction, logger *log.Logger, sid *sid.Sid, jwt *jwt.JWT) *Service {
	return &Service{
		logger: logger,
		sid:    sid,
		jwt:    jwt,
		tm:     tm,
	}
}
