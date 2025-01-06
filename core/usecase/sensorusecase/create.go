package sensorusecase

import (
	"backend/core/domain"
	"backend/core/dto"
)

func (usecase usecase) Create(SensorRequest *dto.CreateSensorRequest) (*domain.Sensor, error) {
	Sensor, err := usecase.repository.Create(SensorRequest)

	if err != nil {
		return nil, err
	}

	return Sensor, nil
}
