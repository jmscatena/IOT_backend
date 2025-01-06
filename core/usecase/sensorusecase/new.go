package sensorusecase

import "backend/core/domain"

type usecase struct {
	repository domain.SensorRepository
}

// New returns contract implementation of SensorUseCase
func New(repository domain.SensorRepository) domain.SensorUseCase {
	return &usecase{
		repository: repository,
	}
}
