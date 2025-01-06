package sensorservice

import "backend/core/domain"

type service struct {
	usecase domain.SensorUseCase
}

// New returns contract implementation of SensorService
func New(usecase domain.SensorUseCase) domain.SensorService {
	return &service{
		usecase: usecase,
	}
}
